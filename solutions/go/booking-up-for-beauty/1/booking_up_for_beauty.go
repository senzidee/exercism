package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
    t, _ := time.Parse(layout, date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    return time.Since(dateToCheck) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    return dateToCheck.Hour() >= 12 && dateToCheck.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    return "You have an appointment on " + dateToCheck.Format("Monday, January 2, 2006") + ", at " + dateToCheck.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
}
