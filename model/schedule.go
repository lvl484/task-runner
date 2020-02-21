package model

import (
	"time"
)

// Schedule consist of all fields required for create schedule
type Schedule struct {
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	Count    int       `json:"count"`
	Interval Duration  `json:"interval,string"`
}

func (s *Schedule) SetDefaultValue() {
	if s.Count == 0 {
		s.Count = 1
	}

	if s.Interval == 0 {
		s.Interval = Duration(1 * time.Second)
	}

	if s.StartAt.IsZero() {
		s.StartAt = time.Now()
	}

	if s.EndAt.IsZero() {
		for i := 0; i < s.Count; i++ {
			s.EndAt = s.EndAt.Add(s.Interval.Duration())
		}
	}
}
