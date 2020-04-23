package calendar

import "encoding/json"

/*title: Calendar Data
First, we need to think about which data needs to be stored by our module.
As we want to display a calendar, we'll need to store a date.
Now thankfully, unlike the Gregorian calendar, Ankh-Morpork's calendar has is very regular – no leap years etc.
So let's implement a simple type that stores a data of the Ankh-Morpork calendar.

| Create the new file inside the `calendar` package.
*/

// Month holds a Discworld month.
type Month int

// constants for each month
const (
	Ick Month = iota
	Offle
	February
	March
	April
	May
	June
	Grune
	August
	Spune
	Sektober
	Ember
	December
)

func (m Month) String() string {
	return [...]string{"Ick", "Offle", "February", "March",
		"April", "May", "June", "Grune", "August", "Spune",
		"Sektober", "Ember", "December"}[m]
}

/*
Here, we create a type for Discworld's months.
Since Go does not have enums, we create a new `int`, some constants, and define its conversion to string.
*/

// UniversityDate stores days since the 1st of Ick, year 0.
type UniversityDate int

func (d UniversityDate) add(daysDelta int) UniversityDate {
	return d + UniversityDate(daysDelta)
}

func (d UniversityDate) year() int {
	// a common year has 400 days.
	if d < 0 {
		return int((d - 399) / 400) // because Go rounds towards zero
	}
	return int(d / 400)
}

func (d UniversityDate) month() Month {
	// Ick has 16 days, all other months have 32 days.
	return Month((d.add(-400*d.year()) + 16) / 32)
}

func (d UniversityDate) dayOfMonth() int {
	t := d.add(-400 * d.year())
	if t < 16 {
		// 1-based, therefore +1
		return int(t + 1)
	}
	return int(t+16)%32 + 1
}

/*
This makes up our type for a Discworld date.
I couldn't find official information about whether a year 0 exists, so we assume so for simplicity's sake.
We use the term `d.add(-400*d.year())` to reach a value between 0 and 399 for month and day calculation, which would not work for negative numbers if we did `d%400`.
*/

// MarshalJSON returns an object containing Day, Month and Year.
func (d UniversityDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Day   int    `json:"day"`
		Month string `json:"month"`
		Year  int    `json:"year"`
	}{
		Day: d.dayOfMonth(), Month: d.month().String(), Year: d.year(),
	})
}

/*
This func customizes marshaling the type to JSON.
We used an anonymous `struct` type that contains day, month and year instead of the total days we store internally.
With the `json:` annotation, we tell the serializer that we want lower-case fields in the resulting JSON.

Generally, QuestScreen serializes data in two ways:

 * Everything being sent to and received from the Web Client is serialized as JSON.
 * Everything being stored to and loaded from the file system is serialized as YAML.

This makes it easy for you to define the way your data is serialized/deserialized depending on the use-case (for YAML storage, you'd define `MarshalYAML`).
The reason we define this marshaler is so that we do not need to replicate the logic calculating year, month and day on the JavaScript side.

This is all we need in this file, let's move on and actually do something with the API.
*/
