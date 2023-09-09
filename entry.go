package main

import (
	"time"
)

type Entry struct {
	Title string    `json:"title"`
	Date  EntryDate `json:"date"`
	Keys  []string  `json:"keys"`
	Body  string    `json:"body"`
}

type EntryDate struct {
	Year     int    `json:"year"`
	Month    string `json:"month"`
	Day      int    `json:"day"`
	Weekday  string `json:"weekday"`
	FullDate string `json:"fullDate"`
	// Formatted string
	Stamp time.Time `json:"stamp"`
}

// func (e *Entry) save() error {
// 	// check if present date corresponds to the entry date, only allow saving when true to prevent alteration of the file after the day has passed.
// 	// also check that entry date is not in the future before saving
// 	//save as a json file, with the day of the month(from the entry date) as the name, in the folders corresponding to the month and year (also gotten from the entry date). Create "title", "date", "keys" and "content" properties with their corresponding values.
// 	return nil
// }
