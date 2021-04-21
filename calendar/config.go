package calendar

import "github.com/QuestScreen/api/config"

/*title: Module configuration
This file contains the data structure that defines the configurable values of
this module.
*/

/*
First, we'll need a configuration object for our module.
Configuration objects hold data which will be configurable on the Web Client's *Configuration* page.

Each field of our configuration's type has to be a pointer to a `config.Item`.
We will use predefined item types here.
It is also possible to define additional configuration items, this is not covered here.

The config item types we are using are:

 * `FontSelect`, which allows the user to select the font with which text is rendered.
   Fonts are taken from the `fonts` directory inside the installation's data folder.

	 It is usually advicable to provide a `FontSelect` setting when you render text.
	 This enables easy customizability for the user.
 * `BackgroundSelect` allows the user to customize the background color.

   This is a recommended setting for every module that renders some kind of background box.
	 It also allows the user to use a texture so that the background is more interesting.

Generally, it is advisable to not hard-code colors or fonts into a plugin, so that the user can customize the appearance to match with the configured look of other plugins.
*/
type calendarConfig struct {
	Font       *config.FontSelect
	Background *config.BackgroundSelect
}
