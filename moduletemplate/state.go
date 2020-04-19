package moduletemplate

import (
	"github.com/QuestScreen/api"
	"gopkg.in/yaml.v3"
)

type state struct {
	// TODO: put the state's data here.
}

// HTTP endpoint (copy & rename if you have more than one)
type endpoint struct {
	*state
}

func newState(input *yaml.Node, ctx api.ServerContext,
	ms api.MessageSender) (api.ModuleState, error) {
	s := new(state)
	// TODO: fill state data
	return s, nil
}

func (s *state) WebView(ctx api.ServerContext) interface{} {
	// TODO: implement
	return nil
}

func (s *state) PersistingView(ctx api.ServerContext) interface{} {
	// TODO: implement
	return nil
}

func (s *state) CreateRendererData() interface{} {
	// TODO: implement
	return nil // this is the data that will be received by Renderer.RebuildState.
}

// PureEndpoint returns the pure endpoint at the given index.
// TODO: can be removed if module has no pure endpoints.
// TODO: IDEndpoint must be added if module has endpoints taking an ID
//       in their URI path.
func (s *state) PureEndpoint(index int) api.ModulePureEndpoint {
	// TODO: modify this if you have more than one endpoints
	if index != 0 {
		panic("Endpoint index out of bounds")
	}
	return endpoint{s}
}

func (e endpoint) Post(payload []byte) (interface{}, interface{},
	api.SendableError) {
	// TODO: implement
	// first value is sent back to client as JSON.
	// second value is sent to Renderer.InitTransition.
	// third value is for errors.
	return nil, nil, nil
}
