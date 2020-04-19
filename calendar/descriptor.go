package calendar

/*title: Module Descriptor
This file contains the module's metadata.
*/

import "github.com/QuestScreen/api"

// Descriptor describes this module.
var Descriptor = api.Module{
	Name:                "Calendar",
	ID:                  "tutorial-calendar",
	ResourceCollections: []api.ResourceSelector{},
	EndpointPaths: []string{
		"", // endpoint with no further path; handles updating the date
	},
	DefaultConfig: &config{Font: &api.SelectableFont{
		FamilyIndex: 0, Size: api.HeadingFont, Style: api.Bold},
		Background: &api.SelectableTexturedBackground{
			Primary:      api.RGBColor{Red: 255, Green: 255, Blue: 255},
			TextureIndex: -1,
		},
	},
	CreateRenderer: newRenderer,
	CreateState:    newState,
}

/*
*Name* is used in the UI to identify the module; it will be prefixed by the plugin name there.
*ID* must be unique amongst all available modules.
It is generally a good idea to use the format `<plugin-name>-<module-name>`.

The empty endpoint path means that our endpoint is reachable at `/state/tutorial-calendar`.
Since we only have one endpoint, we don't need subpaths.

The `DefaultConfig` is the config used when the user does not select something else.
As we can't know which fonts the user has installed, we simply use the first one.
If there are no fonts installed, the user will see an error message in the Web Client and QuestScreen cannot be used, so we don't need to accommodate for that situation.
We want the text large and bold by default.

For the background, we set the primary color to white and the `TextureIndex` to -1 meaning *no texture*.
*/
