package main

import "fmt"

//Months with Days, Hours and Minutes
type Months uint

//Days Hours and Minutes
type Days uint

//Hours and Minutes
type Hours uint

//Minutes calculated
type Minutes uint

//Days calculated from month
func (month Months) Days() Days {
	return Days(month) * 30
}

//Hours calculated from Days
func (days Days) Hours() Hours {
	return Hours(days) * 24
}

//Minutes calculated from Hours
func (hours Hours) Minutes() Minutes {
	return Minutes(hours) * 60
}

func main() {
	var months Months
	months = 30
	fmt.Printf("days in a month %d ", months.Days())
	fmt.Printf("Hours in a month %d ", months.Days().Hours())
	fmt.Printf("Minutes in a month %d ", months.Days().Hours().Minutes())
}
