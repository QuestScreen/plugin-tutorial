package moduletemplate

import "github.com/QuestScreen/api"

// Descriptor describes this module.
var Descriptor = api.Module{
	Name:                "My Module",         // TODO: change
	ID:                  "myplugin-mymodule", // TODO: change
	ResourceCollections: []api.ResourceSelector{
		// TODO: put resource selectors here if you want the module load files
		//       from the file system.
	},
	EndpointPaths: []string{
		"", // for a single HTTP endpoint, see API docs for details.
	},
	DefaultConfig: &config{
		// TODO: fill with default values if config has fields.
	},
	CreateRenderer: newRenderer,
	CreateState:    newState,
}
