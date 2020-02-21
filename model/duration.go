package model

import (
	"fmt"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	a, err := time.ParseDuration(string(b))
	if err != nil {
		return fmt.Errorf("invalid format: %w", err)
	}
	*d = Duration(a)
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v"`, d)), nil
}

func (d Duration) Duration () time.Duration{
	return time.Duration(d)
}
