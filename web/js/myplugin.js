// TODO: change myplugin & mymodule
tmpl.myplugin = {
  mymodule: {
    state: new Template("#tmpl-myplugin-mymodule-state",
        function (state, ctrl) {
      // TODO: render current state into the template.
    })
  }
};

// TODO: change name
class MyModule {
  constructor() {
    // TODO: change id; this must match the Id given in the module's descriptor
    this.id = "myplugin-mymodule";
  }

  ui(app, state) {
    return tmpl.myplugin.mymodule.state.render(state, this);
  }

  // TODO: add controller functions invoked by UI controls.
}

app.registerStateController(new MyModule());