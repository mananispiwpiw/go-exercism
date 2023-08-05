package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var dateTime time.Time
	layoutFormat := "1/2/2006 15:04:05"
	dateTime, _ = time.Parse(layoutFormat, date)
	return dateTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateNow := time.Now()
	layoutFormat := "January 1, 2006 15:04:05"
	dateTime, _ := time.Parse(layoutFormat, date)
	isBefore := dateTime.Before(dateNow)
	return isBefore
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layoutFormat := "Monday, January 1, 2006 15:04:05"
	dateTime, _ := time.Parse(layoutFormat, date)
	if dateTime.Hour() >= 12 && dateTime.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layoutFormat1 := "1/2/2006 15:04:05"
	dateTime, _ := time.Parse(layoutFormat1, date)

	jam := dateTime.Format("15")
	menit := dateTime.Format("04")
	hari := dateTime.Weekday().String()
	tahun := dateTime.Format("2006")
	tanggal := dateTime.Format("2")
	bulan := dateTime.Format("January")

	result := "You have an appointment on " + hari + ", " + bulan + " " + tanggal + ", " + tahun + ", at " + jam + ":" + menit + "."
	return result
	//You have an appointment on Monday, September 19, 1994, at 12:15.
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	nowDate := time.Now()
	var anniversary = time.Date(nowDate.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
