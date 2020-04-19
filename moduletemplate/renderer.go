package moduletemplate

import (
	"time"

	"github.com/QuestScreen/api"
	"github.com/veandco/go-sdl2/sdl"
)

type config struct {
	// TODO: put any config values here
}

// Renderer implements the rendering of the module's state with SDL.
type Renderer struct {
	config *config // holds current merged configuration values
	// add data fields here
}

func newRenderer(backend *sdl.Renderer,
	ms api.MessageSender) (api.ModuleRenderer, error) {
	val := new(Renderer)
	// initialize Renderer fields here
	return val, nil
}

// Descriptor returns the module's metadata
func (*Renderer) Descriptor() *api.Module {
	return &Descriptor
}

// SetConfig sets the merged configuration.
func (r *Renderer) SetConfig(value interface{}) {
	r.config = value.(*config)
}

// InitTransition starts transitioning after user input changed the state.
func (r *Renderer) InitTransition(
	ctx api.RenderContext, data interface{}) time.Duration {
	// TODO: implement
	return 0
}

// TransitionStep advances the transitioning animation.
func (r *Renderer) TransitionStep(
	ctx api.RenderContext, elapsed time.Duration) {
	// TODO: implement
}

// FinishTransition finalizes the transitioning animation.
func (r *Renderer) FinishTransition(ctx api.RenderContext) {
	// TODO: implement
}

// Render renders the current state / animation frame.
func (r *Renderer) Render(ctx api.RenderContext) {
	// TODO: implement
}

// Rebuild rebuilds the state from the given config and optionally data.
func (r *Renderer) Rebuild(
	ctx api.ExtendedRenderContext, data interface{}, configVal interface{}) {
	// TODO: implement
}
