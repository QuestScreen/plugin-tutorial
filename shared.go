package shared

/*title: Shared Data
First, we need to think about what data our module state contains.
The data types used to model this data will be used both at the server and in the web UI, so we need to place them in a common `shared` package.

As we want to display a calendar, we'll need to store a date.
Now thankfully, unlike the Gregorian calendar, Ankh-Morpork's calendar has is very regular – no leap years etc.
So let's implement a simple type that stores a data of the Ankh-Morpork calendar.

Create the new file directly inside `plugin-tutorial`.
Putting it into the `calendar` directory would later lead to the web UI depending on our whole module implementation, which won't work because of dependencies like SDL.
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

// Add adds the given amount of days to the date.
func (d UniversityDate) Add(daysDelta int) UniversityDate {
	return d + UniversityDate(daysDelta)
}

// Year returns the current year (0-based) of the given date.
func (d UniversityDate) Year() int {
	// a common year has 400 days.
	if d < 0 {
		return int((d - 399) / 400) // because Go rounds towards zero
	}
	return int(d / 400)
}

// Month returns the current month of the given date.
func (d UniversityDate) Month() Month {
	// Ick has 16 days, all other months have 32 days.
	return Month((d.Add(-400*d.Year()) + 16) / 32)
}

// DayOfMonth returns the current day within the current month of the given date.
func (d UniversityDate) DayOfMonth() int {
	t := d.Add(-400 * d.Year())
	if t < 16 {
		// 1-based, therefore +1
		return int(t + 1)
	}
	return int(t+16)%32 + 1
}

/*
This makes up our type for a Discworld date.
I couldn't find official information about whether a year 0 exists, so we assume so for simplicity's sake.
We use the term `d.Add(-400*d.Year())` to reach a value between 0 and 399 for month and day calculation, which would not work for negative numbers if we did `d%400`.
*/
