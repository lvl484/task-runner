package model

import (
	"time"
)

// Schedule consist of all fields required for create schedule
type Schedule struct {
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at,omitempty"`
	Count    int       `json:"count,omitempty"`
	Interval Duration  `json:"interval,string"`
}

func (s *Schedule) UntilStartTime() time.Duration {
	return s.StartAt.Sub(time.Now())
}

func (s *Schedule) UntilEndCount() time.Duration {
	return s.StartAt.Add(time.Duration(s.Count) * s.Interval.Duration()).Sub(s.StartAt) - s.Interval.Duration()
}

func (s *Schedule) UntilEndTime() time.Duration {
	return s.EndAt.Sub(time.Now())
}

func (s *Schedule) SetDefaultValue() {
	if s.Interval == 0 {
		s.Interval = Duration(1 * time.Second)
	}

	if s.StartAt.IsZero() {
		s.StartAt = time.Now()
	}
}
