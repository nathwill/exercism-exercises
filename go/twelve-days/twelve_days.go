// Package twelve contains data and methods for handling twelve days of christmas
package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

var lead = "On the %s day of Christmas my true love gave to me, "
var gifts = map[int]string{
	1:  "Partridge in a Pear Tree",
	2:  "Turtle Doves",
	3:  "French Hens",
	4:  "Calling Birds",
	5:  "Gold Rings",
	6:  "Geese-a-Laying",
	7:  "Swans-a-Swimming",
	8:  "Maids-a-Milking",
	9:  "Ladies Dancing",
	10: "Lords-a-Leaping",
	11: "Pipers Piping",
	12: "Drummers Drumming",
}
var days = map[int][2]string{
	1:  {"a", "first"},
	2:  {"two", "second"},
	3:  {"three", "third"},
	4:  {"four", "fourth"},
	5:  {"five", "fifth"},
	6:  {"six", "sixth"},
	7:  {"seven", "seventh"},
	8:  {"eight", "eighth"},
	9:  {"nine", "ninth"},
	10: {"ten", "tenth"},
	11: {"eleven", "eleventh"},
	12: {"twelve", "twelfth"},
}

// Song generates the full Twelve Days of Christmas song
func Song() string {
	var verses []string
	for i := 1; i <= len(days); i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n") + "\n"
}

// Verse returns the verse for a given day.
func Verse(day int) string {
	return fmt.Sprintf(lead, days[day][1]) + giftList(day)
}

// giftList returns the list of gifts for a given day as a string
func giftList(day int) string {
	var gifts []string
	// early exit for the first day
	if day == 1 {
		return giftItem(day) + "."
	}
	// gifts include today and all prior (lesser) days
	for i := day; i > 1; i-- {
		gifts = append(gifts, giftItem(i))
	}
	gifts = append(gifts, "and "+giftItem(1))
	return strings.Join(gifts, ", ") + "."
}

// giftItem returns the gift for a given day as a string
func giftItem(day int) string {
	return fmt.Sprintf("%s %s", days[day][0], gifts[day])
}
