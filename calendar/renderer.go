package calendar

import (
	"fmt"
	"time"

	"github.com/QuestScreen/api/modules"
	"github.com/QuestScreen/api/render"
	"github.com/QuestScreen/api/server"
	shared "github.com/QuestScreen/plugin-tutorial"
)

/*title: Module Renderer
This file contains the code that renders the module to the screen.
*/

// calendarRenderer implements the rendering of the module's state with SDL.
type calendarRenderer struct {
	config         *calendarConfig
	curTex, oldTex render.Image
	cur            shared.UniversityDate
	oldPos         int32
	oldAlpha       uint8
}

/*
This is our renderer, implementing `api.ModuleRenderer`.
Let's discuss the data we put in here:

 * `config` is the current configuration, merged from the several configuration levels (default, base, system, group and scene).
 * `curTex` contains the rendered image of the current date.
   Whenever the date changes, we will render it once into a texture and can then use that texture every time we need to paint the module.
   We do this since we want calls to `Render` be as fast as possible.
 * `oldTex` contains the previous date while we'll use to animate the transition to the new date.
 * `mask` contains a mask tile for the background.
   As you can see with the HeroList and Title modules, the user can configure a secondary color together with a texture that blends it with the primary color.
   We'll use this mask to do that blending, details will be explained later.
 * `cur` is the current date.
   We need to store it because it is possible that the configuration changes without any data change.
   When this happens, we will need to re-render the module, but will not get data sent from the ModuleState.
   In that case, we'll use the data in `cur`.
 * `oldPos` is the position of the old image we'll use during animation.
 * `oldAlpha` is the alpha value of the old image we'll use during animation.
*/

func newRenderer(ctx render.Renderer,
	ms server.MessageSender) (modules.Renderer, error) {
	return &calendarRenderer{}, nil
}

/*
These are trivial funcs we need to implement.
`newRenderer` initializes the Renderer; we don't need to do anything here since everything will be initialized in `Rebuild`.
Consult the diagram in the [Documentation](/plugins/documentation/#The%20Render%20Loop) for details on the order in which ModuleRenderer funcs are called.
*/

func (cr *calendarRenderer) createDateSheet(ctx render.Renderer,
	d shared.UniversityDate) render.Image {
	str := fmt.Sprintf("%d %s %d", d.DayOfMonth(), d.Month(), d.Year())
	strTexture := ctx.RenderText(str, cr.config.Font.Font)
	defer ctx.FreeImage(&strTexture)

	canvas, frame := ctx.CreateCanvas(strTexture.Width+2*ctx.Unit(),
		strTexture.Height+2*ctx.Unit(), cr.config.Background.Background,
		render.East|render.South|render.West)
	frame = frame.Position(strTexture.Width, strTexture.Height, render.Center,
		render.Middle)
	strTexture.Draw(ctx, frame, 255)
	return canvas.Finish()
}

/*
This is a helper function that renders an image of our date.

Of course, we format the date like a sane person would do: *day month year*.
We query the currently selected font face from the context and render our date string to a texture.
This creates a texture `strTexture` that contains our text printed in the given color with transparent background.

Then, we create a *canvas* based on the dimensions of the rendered text.
A canvas redirects all rendering to a texture, which can later be queried with `canvas.Finish()`.
We calculate the inner dimensions with `ctx.Unit()`, which exists to accomodate for different display sizes and dimensions.
This means that our rendered data will occupy the same percentage of width on a FullHD screen as it will on a 4k screen.

`CreateCanvas` optionally renders a background on it, which may use a texture.
We can just give the configured background to the renderer.
A canvas can have borders, which will extend the canvas' size from the dimensions we give.
The borders are specified via flags.
Since we will anchor our date at the top edge of the screen, we create borders for the other three directions.

`CreateCanvas` also returns a frame which just describes the dimensions of the canvas.
Being a `Rectangle`, it provides several functions to position sub-rectangles in it.
One of those is `Position`, which we use to center a rectangle of the size of our text horizontally and vertically centered in the canvas.

Now we draw the rendered text into the remaining frame inside the canvas.
We have the possibility to use blending by setting an alpha value lower than 255, but we do not need this for our purposes.

Finally, we finish the canvas, which creates the texture containing our rendered text on the user-chosen background.
This is what we return.
*/

// Rebuild rebuilds the state from the given config and optionally data.
func (cr *calendarRenderer) Rebuild(
	ctx render.Renderer, data interface{}, configVal interface{}) {
	cr.config = configVal.(*calendarConfig)
	if data != nil {
		cr.cur = data.(shared.UniversityDate)
	}
	ctx.FreeImage(&cr.curTex)
	cr.curTex = cr.createDateSheet(ctx, cr.cur)
}

/*
In this function, we need to update the current image based on given configuration value and optionally state data.
The Configuration value will always be given, but state data may be `nil`.
If `data` is not `nil`, it has been generated by `ModuleState`'s `CreateRendererData`.

Since our images are OpenGL textures, they are not automatically garbage-collected.
We need to be careful to always destroy them with `FreeImage` before creating a new image.
We can call this without checking if the image actually exists because it does nothing on empty images.

---

Before we implement the animation now, let's discuss how it should look like:

When a new date is set, we want to rip of the old data like we'd do with a calendar sheet.
This means that we'll render the new date at its final position, and over it the old date that falls down and fades away.

This is a pretty easy animation; we only need the images of the old and new date, and update the position and transparency of the old image with each step.
*/

// InitTransition starts transitioning after user input
// changed the state.
func (cr *calendarRenderer) InitTransition(
	ctx render.Renderer, data interface{}) time.Duration {
	cr.oldTex = cr.curTex
	cr.cur = data.(shared.UniversityDate)
	cr.curTex = cr.createDateSheet(ctx, cr.cur)
	cr.oldPos = 0
	return time.Second / 2
}

/*
This func receives the data returned by our endpoint.
We move the old image to `oldTex` since we still need it for the animation.
Then we draw the new image.
The initial animation state will be the old texture being completely visible and at the original position.
We return the time span used for the animation.
*/

// TransitionStep advances the transitioning animation.
func (cr *calendarRenderer) TransitionStep(
	ctx render.Renderer, elapsed time.Duration) {
	pos := render.TransitionCurve{Duration: time.Second / 2}.Cubic(elapsed)
	cr.oldAlpha = uint8((1.0 - pos) * 255)
	cr.oldPos = int32(pos * float32(cr.oldTex.Height) * 3)
}

/*
When advancing the animation, we use a `TransitionCurve`, which implements a function going from `0.0` at the beginning to `1.0` at the end of the animation.
Generally, a linear progression looks very artificial.
The `Cubic` curve we use starts slow, speeds up, and decelerates at the end.

We set `oldAlpha` to facilitate fading, and `oldPos` defines how far down the old image is.
We use the image's height for defining how far it moves.
*/

// FinishTransition finalizes the transitioning animation.
func (cr *calendarRenderer) FinishTransition(ctx render.Renderer) {
	ctx.FreeImage(&cr.oldTex)
	cr.oldAlpha = 0
}

/*
At the end of the animation, we destroy the old texture.
We do not need to reset `oldPos` since that is not used outside of animation.
It will be re-initialized when a new animation starts.
*/

// Render renders the current state / animation frame.
func (cr *calendarRenderer) Render(ctx render.Renderer) {
	frame := ctx.OutputSize()

	_, frame = frame.Carve(render.East, 5*ctx.Unit())

	cr.curTex.Draw(ctx,
		frame.Position(cr.curTex.Width, cr.curTex.Height, render.Right, render.Top),
		255-cr.oldAlpha)
	if !cr.oldTex.IsEmpty() {
		_, frame = frame.Carve(render.North, cr.oldPos)
		cr.oldTex.Draw(ctx,
			frame.Position(cr.oldTex.Width, cr.oldTex.Height, render.Right, render.Top),
			cr.oldAlpha)
	}
}

/*
Finally, rendering.
We render the calender to the upper right corner, with a distance of 5 units from the right edge.
This is done by first carving out the 5 units from the right and then positioning the text's rectangle at the top right of the remaining frame.

If `cr.oldTex` contains an image, we're currently animating so we need to render the old date as well.
For this, we use the value `cr.oldPos` calculated in our `TransitionStep` to offset the old texture from the screen top.

This wraps up the code for rendering.
*/
