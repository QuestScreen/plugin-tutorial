package calendar

/*title: Module Renderer
We now need to implement `api.ModuleRenderer`.
This is the interface for rendering our module to the screen.
For this part, we use the SDL2 library but we'll be doing nothing too complex, since the plugin API provides some helper functions for it.
*/

import (
	"fmt"
	"time"

	"github.com/QuestScreen/api"
	"github.com/veandco/go-sdl2/sdl"
)

type config struct {
	Font       *api.SelectableFont
	Background *api.SelectableTexturedBackground
}

/*
First, we'll need a configuration object for our module.
Configuration objects hold data which will be configurable on the Web Client's *Configuration* page.

Each field of our configuration's type has to be a pointer to an `api.ConfigItem`.
We will use predefined item types here; if you want to define custom config item types, you will not only need to implement `api.ConfigItem`, but also provide JavaScript code to the WebClient for editing those item types.

The config item types we're using are:

 * `SelectableFont`, which allows the user to select the font with which text is rendered from the fonts available in the current installation.
	 It is advicable to use this if you render any text; this enables the user to use any font they like.
 * `SelectableTexturedBackground` allows the user to customize the background color.

Generally, it is advisable to not hard-code colors or fonts into a plugin, so that the user can customize the appearance to match with the configured look of other plugins.
*/

// Renderer implements the rendering of the module's state with SDL.
type Renderer struct {
	config         *config
	curTex, oldTex *sdl.Texture
	mask           *sdl.Texture
	cur            UniversityDate
	oldPos         int32
}

/*
This is our renderer, implementing `api.ModuleRenderer`.
Let's discuss the data we put in here:

 * `config` is the current configuration, merged from the several configuration levels (default, base, system, group and scene).
 * `curTex` contains the rendered image of the current date.
	 Whenever the date changes, we will render it once into a texture and can then use that texture every time we need to paint the module.
	 We do this since we want calls to `Render` be as fast as possible.
 * `oldTex` contains the previous date while we'll do animation to transition to the new date.
 * `mask` contains a mask tile for the background.
	 As you can see with the HeroList and Title modules, the user can configure a secondary color together with a texture that blends it with the primary color.
	 We'll use this mask to do that blending, details will be explained later.
 * `cur` is the current date.
	 We need to store it because it is possible that the configuration changes without any data change.
	 When this happens, we will need to re-render the module, but will not get data sent from the ModuleState.
	 In that case, we'll use the data in `cur`.
 * `oldPos` is the position of the old image we'll use during animation.
*/

func newRenderer(backend *sdl.Renderer,
	ms api.MessageSender) (api.ModuleRenderer, error) {
	return &Renderer{}, nil
}

// Descriptor returns the module's metadata
func (*Renderer) Descriptor() *api.Module {
	return &Descriptor
}

/*
These are trivial funcs we need to implement.
`newRenderer` initializes the Renderer; we don't need to do anything here since everything will be initialized in `Rebuild`.
Consult the diagram in the [Documentation](/plugins/documentation/#The%20Render%20Loop) for details on the order in which ModuleRenderer funcs are called.
*/

func (r *Renderer) createDateSheet(ctx api.RenderContext,
	d UniversityDate) *sdl.Texture {
	str := fmt.Sprintf("%d %s %d", d.dayOfMonth(), d.month(), d.year())
	face := ctx.Font(
		r.config.Font.FamilyIndex, r.config.Font.Style, r.config.Font.Size)
	strTexture := ctx.TextToTexture(str, face, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	_, _, strWidth, strHeight, _ := strTexture.Query()
	bgColor := r.config.Background.Primary.WithAlpha(255)
	canvas := ctx.CreateCanvas(strWidth+2*ctx.Unit(), strHeight+2*ctx.Unit(),
		&bgColor, r.mask, api.East|api.South|api.West)
	ctx.Renderer().Copy(strTexture, nil, &sdl.Rect{
		X: 2 * ctx.Unit(), Y: ctx.Unit(), W: strWidth, H: strHeight})
	return canvas.Finish()
}

/*
This is a helper function that renders an image of our date.

Of course, we format the date like a sane person would do: *day month year*.
We query the currently selected font face from the context and render our date string to a texture.
This creates a texture that contains our text printed in the given color with transparent background.

Then, we create a *canvas* based on the dimensions of the rendered text.
A canvas redirects all rendering to a texture, which can later be queried with `canvas.Finish()`.

We calculate the inner dimensions with `ctx.Unit()`, which exists to accomodate for different display sizes and dimensions.
This means that our rendered data will occupy the same percentage of width on a FullHD screen as it will on a 4k screen.

`CreateCanvas` optionally renders a background color and mask on it.
We give the selected color and the current mask to it so that it does that for us.

A canvas can have borders, which will extend the canvas' size from the dimensions we give.
The borders are specified via flags.
Since we will anchor our date at the top edge of the screen, we create borders for the other three directions.

Now we copy the rendered text with the renderer, offsetting it so that it is centered.
To calculate the correct x offset, we need to account for both the inner padding we included in the canvas size, and the outer border the canvas added based on the direction flags.

Finally, we finish the canvas and return the result.
*/

// Rebuild rebuilds the state from the given config and optionally data.
func (r *Renderer) Rebuild(
	ctx api.ExtendedRenderContext, data interface{}, configVal interface{}) {
	r.config = configVal.(*config)
	ctx.UpdateMask(&r.mask, *r.config.Background)
	if data != nil {
		r.cur = data.(UniversityDate)
	}
	if r.curTex != nil {
		r.curTex.Destroy()
	}
	r.curTex = r.createDateSheet(ctx, r.cur)
}

/*
In this function, we need to update the current image based on given configuration value and optionally state data.
The Configuration value will always be given, but state data may be `nil`.
If `data` is not `nil`, it has been generated by `ModuleState`'s `CreateRendererData`.

The API provides us with a function to regenerate the mask based on the current configuration, so we do not need to care about the details.

Since SDL is a C library, it does not automatically delete data.
We need to be careful to always destroy SDL objects we do not need anymore *before* we assign a new value to it.

---

Before we implement the animation now, let's discuss how it should look like:

When a new date is set, we want to rip of the old data like we'd do with a calendar sheet.
This means that we'll render the new date at its final position, and over it the old date that falls down and fades away.

This is a pretty easy animation; we only need the images of the old and new date, and update the position and transparency of the old image with each step.
*/

// InitTransition starts transitioning after user input changed the state.
func (r *Renderer) InitTransition(
	ctx api.RenderContext, data interface{}) time.Duration {
	r.oldTex = r.curTex
	r.oldTex.SetBlendMode(sdl.BLENDMODE_BLEND)
	r.cur = data.(UniversityDate)
	r.curTex = r.createDateSheet(ctx, r.cur)
	r.oldPos = 0
	return time.Second / 2
}

/*
This func receives the data returned by our endpoint.
For the old texture to fade out, we need to activite blending on it.
The initial animation state will be the old texture being completely visible and at the original position.
We return the time span used for the animation.
*/

// TransitionStep advances the transitioning animation.
func (r *Renderer) TransitionStep(
	ctx api.RenderContext, elapsed time.Duration) {
	pos := api.TransitionCurve{Duration: time.Second / 2}.Cubic(elapsed)
	r.oldTex.SetAlphaMod(uint8((1.0 - pos) * 255))
	_, _, _, oldHeight, _ := r.oldTex.Query()
	r.oldPos = int32(pos * float32(oldHeight) * 3)
}

/*
When advancing the animation, we use a `TransitionCurve`, which implements a function going from `0.0` at the beginning to `1.0` at the end of the animation.
Generally, a linear progression looks very artificial.
The `Cubic` curve we use starts slow, speeds up, and decelerates at the end.

We set the texture's alpha mod to facilitate fading, and the `oldPos` defines how far down the old image is.
We use the image's height for defining how far it moves.
*/

// FinishTransition finalizes the transitioning animation.
func (r *Renderer) FinishTransition(ctx api.RenderContext) {
	r.oldTex.Destroy()
	r.oldTex = nil
}

/*
At the end of the animation, we destroy the old texture.
We do not need to reset `oldPos` since that is not used outside of animation.
*/

// Render renders the current state / animation frame.
func (r *Renderer) Render(ctx api.RenderContext) {
	sr := ctx.Renderer()
	screenWidth, _, _ := sr.GetOutputSize()
	_, _, curWidth, curHeight, _ := r.curTex.Query()
	sr.Copy(r.curTex, nil, &sdl.Rect{X: screenWidth - curWidth - 5*ctx.Unit(),
		Y: 0, W: curWidth, H: curHeight})
	if r.oldTex != nil {
		_, _, oldWidth, oldHeight, _ := r.oldTex.Query()
		sr.Copy(r.oldTex, nil, &sdl.Rect{X: screenWidth - oldWidth - 5*ctx.Unit(),
			Y: r.oldPos, W: oldWidth, H: oldHeight})
	}
}

/*
Finally, rendering.
We render the calender to the upper right corner, with a distance of 5 units from the right edge.
If `r.oldTex` is not nil, we're currently animating so we need to render the old date as well.
Fading and position assignment has already been handled by `TransitionStep`.

This wraps up the code for rendering.
*/
