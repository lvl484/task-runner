package model

import (
	"time"
)

var Never time.Time

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

func (s *Schedule) Next() time.Time {
	if s.Count != nil {
		*s.Count--
		if *s.Count < 0 {
			return Never
		}
	}

	now := time.Now()

	if s.StartAt != nil {
		now = *s.StartAt
	}

	if s.Year != nil {
		now = now.AddDate(*s.Year - now.Year(), 0, 0)
	}

	if s.Month != nil && IsValidMonth(*s.Month){
		m := *s.Month - int(now.Month())
		y := 0
		if m < 0 {
			y = 1
		}
		now = now.AddDate(y, m, 0)
	}

	if s.Day != nil && IsValidDay(*s.Day){
		d := *s.Day - now.Day()
		m := 0
		if d < 0 {
			m = 1
		}
		now = now.AddDate(0, m, d)
	}

	if s.Hour != nil && IsValidHours(*s.Hour){
		h := *s.Hour - now.Hour()
		if h < 0 {
			h = 24 + h
		}
		now = now.Add(time.Duration(h) * time.Hour)
	}

	if s.Minute != nil && IsValidMinutes(*s.Minute){
		m := *s.Minute - now.Minute()
		if m < 0 {
			m = 60 + m
		}
		now = now.Add(time.Duration(m) * time.Minute)
	}

	if s.Second != nil && IsValidSeconds(*s.Second){
		s := *s.Second - now.Second()
		if s < 0 {
			s = 60 + s
		}
		now = now.Add(time.Duration(s) * time.Second)
	}

	if s.Week != nil {
		w := now.Day() / 7
		dw := *s.Week - w
		if dw < 0 {
			dw += w
		}
		now = now.AddDate(0, 0, dw*7)
	}

	if s.Weekday != nil && IsValidWeekDay(*s.Weekday){
		wd := *s.Weekday - int(now.Weekday())
		if wd < 0 {
			wd = 7 + wd
		}
		now = now.AddDate(0, 0, int(wd))
	}

	if s.Interval != nil {
		now = now.Add(*s.Interval)
	}

	if s.EndAt != nil && now.After(*s.EndAt) {
		return Never
	}

	return now
}
