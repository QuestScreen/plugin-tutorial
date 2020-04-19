/*title: Web UI Controller
Here we'll implement the controller for our UI.
Generally, the controller needs to provide functionality to build the UI for our module's state, and also handle all its actions.
*/

const tutorialCalendarSteps = [-10, -3, -1, 1, 3, 10];

/*
Pretty straightforward: These are the step sizes for each button we want to generate.
*/

tmpl.tutorial = {
  calendar: {
    row: new Template("#tmpl-tutorial-calendar-row",
        function (ctrl, handler) {
      for (const [index, button] of this.querySelectorAll("button").entries()) {
        button.addEventListener("click",
            () => handler.call(ctrl, tutorialCalendarSteps[index]));
      }
    }),
    state: new Template("#tmpl-tutorial-calendar-state",
        function (ctrl) {
      const tbody = this.querySelector("tbody");

      const dayRow = tmpl.tutorial.calendar.row.render(
          ctrl, ctrl.changeDay);
      tbody.appendChild(dayRow);

      const monthRow = tmpl.tutorial.calendar.row.render(
          ctrl, ctrl.changeMonth);
      tbody.appendChild(monthRow);

      const yearRow = tmpl.tutorial.calendar.row.render(
          ctrl, ctrl.changeYear);
      tbody.appendChild(yearRow);

      [ctrl.daySpan, ctrl.monthSpan, ctrl.yearSpan] =
          Array.from(tbody.querySelectorAll("span"));
    })
  }
};

/*
`tmpl` is a global object that's used for all templates.
It's not required, but recommended to use it.
Inside, we add an object for our plugin, and in there, our module's template renderers.
A template renderer is an instance of the class `Template`, which implements a method `render`.
`render` copies the children of the HTML `<template>` element into the current DOM and then calls the user-provided renderer function on it.

Those user-provided renderer functions are where we inject values into the template.
`state` is the main template and adds three instances of `row` to the table; each `row` links its buttons to the relevant handler of the controller in `ctrl`.

To be able to write the current day, month and year value easily, we register a reference to the `span` elements containing those values with the controller.
*/

class TutorialCalendar {
  constructor() {
    this.id = "tutorial-calendar";
  }

  updateData(state) {
    this.daySpan.textContent = state.day;
    this.monthSpan.textContent = state.month;
    this.yearSpan.textContent = state.year;
    this.state = state;
  }

  ui(app, state) {
    // this will assign daySpan, monthSpan and yearSpan.
    let ret = tmpl.tutorial.calendar.state.render(this);
    this.updateData(state);
    return ret;
  }

  async changeDay(amount) {
    this.updateData(await App.fetch("state/tutorial-calendar", "POST", amount));
  }

  async changeMonth(amount) {
    this.updateData(
      await App.fetch("state/tutorial-calendar", "POST", amount * 32));
  }

  async changeYear(amount) {
    this.updateData(
      await App.fetch("state/tutorial-calendar", "POST", amount * 400));
  }
}

/*
This is our controller class.
The helper function `updateData` displays the current date in the UI.
It also stores the `state` in the controller.
Why would we need it?
Because `changeMonth` currently doesn't work correctly since *Ick* has only 16 days and we should act correctly on the user skipping from, to or over that month.
This is left as an exercise to the reader.

With `App.fetch`, we query our module endpoint, which takes the number of days we want to modify as parameter.
It returns the new date, and we update the UI accordingly.
Generally, the update logic on data objects should not be reproduced in the client if not absolutely necessary.
Updating the actual values based on the response ensures that the client will not get out of sync with the server, e.g. when the request fails due to temporary network problems.
*/

app.registerStateController(new TutorialCalendar());

/*
This registers our controller with the Web Client.
The field `id` is used to bind it to scenes where a module with that `id` is used.
*/