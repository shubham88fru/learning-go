package main

import (
	"fmt"
	"time"
)

func main() {

	// current local time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC) //2024-07-04 12:00:00 +0000 UTC
	fmt.Println("Specific Time: ", specificTime)

	// format time
	t := time.Now()
	fmt.Println("Fromatted time: ", t.Format("2006-01-02 15:04:05")) // Go uses this specific date to format/parse time

	// time operations
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	// timezone conversion
	loc, _ := time.LoadLocation("America/New_York") // Load the New York timezone
	newYorkTime := time.Now().In(loc)
	fmt.Println("New York Time: ", newYorkTime)
}
