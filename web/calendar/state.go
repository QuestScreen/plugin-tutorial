package calendar

import (
	"encoding/json"

	"github.com/QuestScreen/api/web"
	"github.com/QuestScreen/api/web/modules"
	shared "github.com/QuestScreen/plugin-tutorial"
)

/*title: State UI Implementation
This file contains the code needed for our state UI to function.
*/

type RowKind int

const (
	DayRow RowKind = iota
	MonthRow
	YearRow
)

func (rk RowKind) Caption() string {
	switch rk {
	case DayRow:
		return "Days"
	case MonthRow:
		return "Months"
	default:
		return "Years"
	}
}

/*
First, we need to define the `RowKind` type we used in our HTML components.
This will be seen in the components since the `.askew` file is in the same package.

At this point, you'll want to save the file and run

    askew

in the main directory of the plugin.
This will generate a file `ui.askew.go` which will define the Go types of our
components which we will use in the code that follows.
*/

func NewState(data json.RawMessage, srv web.Server) (modules.State, error) {
	var values shared.UniversityDate

	if err := json.Unmarshal(data, &values); err != nil {
		return nil, err
	}

	ret := &calendarUI{}
	ret.askewInit(srv)
	ret.update(values)

	return ret, nil
}

/*
This is our state constructor function.
It receives a chunk of JSON data via `data`, which is what the server sent us as
current state of the calendar module.

We deserialize this data and initialize our UI component with the resulting
values.
`askewInit` does internal initialization and has the parameter we declared on
our component.
The `update` function loads the received values into our UI, it is defined below.
*/

func (o *calendarUI) update(values shared.UniversityDate) {
	o.days.value.Set(values.DayOfMonth())
	o.months.value.Set(values.Month().String())
	o.years.value.Set(values.Year())
}

/*
Here we access each row (`days`, `months`, `years`) and set their caption to
the current value.
*/

func (o *calendarUI) step(kind RowKind, amount int) {
	switch kind {
	case MonthRow:
		amount *= 32
	case YearRow:
		amount *= 400
	}

	var values shared.UniversityDate

	o.srv.Fetch(web.Post, "", amount, &values)
	o.update(values)
}

/*
Finally, the callback from the buttons.
This will call the server endpoint we defined via `o.srv.Fetch`.
We use the POST method and an empty subpath just as we defined for that endpoint.
We send the amount of days we want to step, and receive the updated values.
With those, we update the UI.
Ta-dah!
*/
