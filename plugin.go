package main

/*title: Finishing the Plugin
Now the module is ready, and we'll finish the plugin metadata.
*/

import (
	"github.com/QuestScreen/api"
	"github.com/QuestScreen/plugin-tutorial/calendar"  // change this
	"github.com/QuestScreen/plugin-tutorial/generated" // and this
)

/*
You need to modify the import paths for `calendar` and `generated` to match the actual path of your plugin in `$GOPATH/src`.
`generated` contains our web files as Go source.
To generate them, you'll need [go-bindata][1] and execute `make generated/data.go` in the plugin's folder.
*/

// QSPlugin is the plugin's descriptor.
var QSPlugin = api.Plugin{
	Name: "Discworld Tutorial",
	Modules: []*api.Module{
		&calendar.Descriptor,
	},
	AdditionalJS:   generated.MustAsset("web/js/controllers.js"),
	AdditionalHTML: generated.MustAsset("web/html/templates.html"),
	AdditionalCSS:  nil,

	/*
	   This is our plugin's descriptor.
	   Its name *must* be `QSPlugin` because that's the name the Core will search for when importing a module.
	   All our web content goes into the `Additional*` fields.
	   If split the code into multiple files, you must join them here into a single byte array.
	   If you changed the file names of the web files, change them here and also update them in the `Makefile`.
	*/

	SystemTemplates: []api.SystemTemplate{
		{
			Name: "Discworld", ID: "tutorial-discworld",
			Config: []byte("name: Discworld"),
		},
	},

	/*
	   Here we define that the plugin requires a system named **Discworld**.
	   This means that if a system with the given ID does not exist at app startup, it will be created with the given config.
	   In the Web UI, that system will not be deletable.
	   The `Config` contains a YAML representation of the system's default configuration.
	   You can use it to define a default look of the system provided by your plugin by also providing configuration for the base plugins.
	   The best way to generate the config string is to load the plugin with a minimal config (as shown), set the desired values via the Web UI, and then copying the contents of the stored `config.yaml` file.
	*/

	GroupTemplates: []api.GroupTemplate{
		{
			Name: "Discworld", Description: "Default Discworld group",
			Config: []byte("name: Discworld\nsystem: tutorial-discworld"),
			Scenes: []api.SceneTmplRef{
				{Name: "Main", TmplIndex: 0},
			},
		},
	},

	/*
	   Next, we're defining a group template, so that users can quickly create a new group using our system & module.
	   In `Config`, we reference the ID of the system config.
	   In `Scenes`, we reference a scene template defined below, via its index.
	   A group template must always refer to at least one scene template, since every group must have at least one scene.
	*/

	SceneTemplates: []api.SceneTemplate{
		{
			Name: "DiscworldMain", Description: "A scene with base modules plus Discworld calendar.",
			Config: []byte(`name: BaseMain
modules:
  background:
    enabled: true
  herolist:
    enabled: true
  overlays:
    enabled: true
  title:
    enabled: true
  tutorial-calendar:
    enabled: true`),
		},
	},
}

/*
Finally, the scene template.
Here, we define a scene that enables all default modules plus our calendar module.
You can enable any modules you know the ID of, so you can reference modules of other plugins here.
There will be a warning if a module referenced by a scene template is not available, but the template will still be usable.
This means that you can have weak dependencies on other plugins.

Be very careful with inline YAML, as it must be indented with spaces, not tabs.
*/

// required to compile; although never called
func main() {}

/*
This function must exist even though we compile the plugin with `-buildmode=plugin`, which causes it to never being called.

 [1]: https://github.com/go-bindata/go-bindata
*/
