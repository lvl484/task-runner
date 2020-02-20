package model

import (
	"time"
)

var Never time.Time

const (
	DayHours = 24
	Minutes = 60
	Seconds = 60
	WeekDays = 7
)

// Schedule consist of all fields required for create schedule
type Schedule struct {
	StartAt  *time.Time     `json:"start_at"`
	EndAt    *time.Time     `json:"end_at"`
	Count    *int           `json:"count"`
	Interval *time.Duration `json:"interval"`

	Year    *int `json:"year"`
	Month   *int `json:"month"`
	Week    *int `json:"week"`
	Weekday *int `json:"weekday"`
	Day     *int `json:"day"`
	Hour    *int `json:"hour"`
	Minute  *int `json:"minute"`
	Second  *int `json:"second"`
}

// Next() returns time when task will be run next time
func (s *Schedule) Next() time.Time {
	if s.Count != nil {
		*s.Count--
		if *s.Count < 0 {
			return Never
		}
	}

	now := time.Now()

	// if StartAt is specified, set var now as it
	if s.StartAt != nil {
		now = *s.StartAt
	}

	// if Year is specified, add to now difference of years
	if s.Year != nil {
		now = now.AddDate(*s.Year - now.Year(), 0, 0)
	}

	// if month is specified, checks it, and set m as difference of months
	// and set y (means Year) as 1 if m is negative (m < 0)
	// add to now y (Year) and m (Month)
	if s.Month != nil && IsValidMonth(*s.Month){
		m := *s.Month - int(now.Month())
		y := 0
		if m < 0 {
			y = 1
		}
		now = now.AddDate(y, m, 0)
	}

	// if day is specified, checks it, and set d as difference of days
	// and set m (means Month) as 1 if d is negative (d < 0)
	// add to now m (Month) and d (Day)
	if s.Day != nil && IsValidDay(*s.Day){
		d := *s.Day - now.Day()
		m := 0
		if d < 0 {
			m = 1
		}
		now = now.AddDate(0, m, d)
	}

	// if hours are specified, checks them, and set h as difference of hours
	// and set h (means Hour) as sum of 24 and h if h is negative (h < 0)
	// add to now duration in hours
	if s.Hour != nil && IsValidHours(*s.Hour){
		h := *s.Hour - now.Hour()
		if h < 0 {
			h = DayHours + h
		}
		now = now.Add(time.Duration(h) * time.Hour)
	}

	// if minute are specified, checks them, and set m as difference of minutes
	// and set m (means Minute) as sum of 60 and m if m is negative (m < 0)
	// add to now duration in minutes
	if s.Minute != nil && IsValidMinutes(*s.Minute){
		m := *s.Minute - now.Minute()
		if m < 0 {
			m = Minutes + m
		}
		now = now.Add(time.Duration(m) * time.Minute)
	}

	// if seconds are specified, checks them, and set s as difference of seconds
	// and set s (means Seconds) as sum of 60 and s if s is negative (s < 0)
	// add to now duration in seconds
	if s.Second != nil && IsValidSeconds(*s.Second){
		s := *s.Second - now.Second()
		if s < 0 {
			s = Seconds + s
		}
		now = now.Add(time.Duration(s) * time.Second)
	}

	// if week is specified, checks it, calculate week number, and set dw as difference of weeks
	// and set dw (means Week) as sum of dw and w if dw is negative (dw < 0)
	// add to now dw * 7 (count of required weeks)
	if s.Week != nil {
		w := now.Day() / WeekDays
		dw := *s.Week - w
		if dw < 0 {
			dw += w
		}
		now = now.AddDate(0, 0, dw*WeekDays)
	}

	// if weekday is specified, checks it, and set wd as difference of week days
	// and set wd (means WeekDay) as sum of 7 and wd if wd is negative (wd < 0)
	// add to now wd (required week days)
	if s.Weekday != nil && IsValidWeekDay(*s.Weekday){
		wd := *s.Weekday - int(now.Weekday())
		if wd < 0 {
			wd = WeekDays + wd
		}
		now = now.AddDate(0, 0, int(wd))
	}

	// if interval is specified, checks it
	// add to now interval
	if s.Interval != nil {
		now = now.Add(*s.Interval)
	}

	// if EndAt specified, compares it with task execution time
	// and return Never if task execution time is After EndAt
	if s.EndAt != nil && now.After(*s.EndAt) {
		return Never
	}

	return now
}
