package calendar

/*title: Module State
In this file, we'll implement `api.ModuleState`.
A module's state will be written to and loaded from the file system and has interfaces to both the module's renderer and the Web Client.
*/

import (
	"github.com/QuestScreen/api/comms"
	"github.com/QuestScreen/api/modules"
	"github.com/QuestScreen/api/server"
	shared "github.com/QuestScreen/plugin-tutorial"
	"gopkg.in/yaml.v3"
)

type state struct {
	Date shared.UniversityDate
}

/*
Our state only holds a UniversityDate.
The `Date` field must be publicly visible so that it can be serialized properly.
*/

type endpoint struct {
	*state
}

/*
The endpoint is the object that is handling requests coming from the Web Client via HTTP.
We need an endpoint object (instead of defining its methods directly on `state`) since we can have more than one endpoint.
*/

func newState(input *yaml.Node, ctx server.Context,
	ms server.MessageSender) (modules.State, error) {
	s := new(state)
	if input == nil {
		return s, nil
	}
	if err := input.Decode(&s.Date); err != nil {
		ms.Warning("unable to load UniversityDate: " + err.Error())
		s.Date = shared.UniversityDate(0)
	}
	return s, nil
}

/*
This is the constructor that is used for creating a `state` from YAML input.
As discussed previously, YAML is the file format all persistent data is stored in.
The input is a subtree of the whole state which also contains other modules.
You do not need to know details about YAML, just use `Decode`.

`input` may be **`nil`** if the currently stored state has no information about our module.
This will always be the case after adding a module to a scene or loading a new group the first time, so we need to deal with it.
Here, we just return a state with the default value (which will be 0, corresponding to 1st of Ick, year 0).

An error during decoding means that the input data is corrupted.
If that is the case, we issue a warning and load the default value.
Returning an error from the module constructor will halt the main app, so don't do it as long as you can load some default value!
*/

func (s *state) Send(ctx server.Context) interface{} {
	return s.Date
}

func (s *state) Persist(ctx server.Context) interface{} {
	return s.Date
}

/*
Now come the serialization functions.
`WebView` returns the data that should be serialized to JSON and send to the web client.
The caller will use JSON serialization on the returned value, which in turn will use `UniversityDate`'s `MarshalJSON` method.

In `PersistingView`, we need to give the same data we `Decode` the input to in the constructor.
This is the data that will be written to the scene state on the file system.
*/

func (s *state) CreateRendererData(ctx server.Context) interface{} {
	return s.Date
}

/*
This function defines the data we send to the renderer so that it can rebuild its state (e.g. when the group is loaded or the scene changes).
As the renderer runs in another thread, it has its own state and cannot access the `state` object.

The returned value must not contain a pointer to data owned by `state` for thread safety as it will be received by the renderer in another thread.
`Date` neither is nor contains a pointer, so we are safe here.
*/

func (s *state) PureEndpoint(index int) modules.PureEndpoint {
	if index != 0 {
		panic("Endpoint index out of bounds")
	}
	return endpoint{s}
}

/*
This function creates our endpoint and implements `api.PureEndpointProvider`.
The module's descriptor will later describe how many and what kind of endpoints a module has, which in turn leads to calls to this function.
Since we only have one endpoint, we can assume that index is always `0`.
*/

func (e endpoint) Post(payload []byte) (interface{}, interface{},
	server.Error) {
	var daysDelta int
	if err := comms.ReceiveData(payload, &daysDelta); err != nil {
		return nil, nil, &server.BadRequest{Inner: err, Message: "received invalid data"}
	}
	e.state.Date = e.state.Date.Add(daysDelta)

	// first value is sent back to client as JSON.
	// second value is sent to Renderer.InitTransition.
	return e.state.Date, e.state.Date, nil
}

/*
Finally, this is our endpoint implementation.
We receive the delta (in days) we want to change, and simply apply it to our date.
`api.ReceiveData` is a helper function that uses JSON unmarshaling, wraps any error into an `api.SendableError`, and can do some additional validation (which we do not need here).
We send the full date back to the Web Client and also to the Renderer.

Generally, a call to an endpoint might lead to a smaller change that does not update the whole data.
For example, think about when you want to hide just one of your heroes with the herolist plugin.
In such a case, we would not send the whole data to the renderer, but a data object that identifies the change.
This way, we can animate small changes (one hero fading out) while keeping other stuff intact.

In our case, to keep things simple, we only want the old date to fade out and the new one to replace it, so there's no point in sending a smaller data package, and thus we just send the whole data.

This is everything we need to do in order to implement `api.ModuleState`.
*/
