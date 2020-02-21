package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Schedule consist of all fields required for create schedule
type Schedule struct {
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	Count    int       `json:"count"`
	Interval Duration  `json:"interval,string"`
	// storing all times when each task will be run
	ExecutionsTimes []time.Time  // TODO: delete in future
}

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	schedule := &Schedule{}

	// Decoding json into object schedule
	err := json.NewDecoder(r.Body).Decode(schedule)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	// Set default values if some fields were not be specified
	if schedule.Count == 0 {
		schedule.Count = 1
	}
	// Set Interval as 1s if Interval field was not be specified
	if schedule.Interval == 0 {
		schedule.Interval = Duration(1 * time.Second)
	}
	// Set StartAt as current time if StartAt was not be specified
	if schedule.StartAt.IsZero() {
		schedule.StartAt = time.Now()
	}
	// Set default EndAt as time in format RFC3339 ("0001-01-01 00:00:00 +0000 UTC")
	if schedule.EndAt.IsZero() {
		schedule.EndAt = schedule.StartAt.Add(schedule.Interval.Duration())
		// Check Count. If it is specified to make EndAt
		if schedule.Count != 1 {
			// increasing EndAt time according to Count
			for i := 0; i < schedule.Count-1; i++ {
				schedule.EndAt = schedule.EndAt.Add(schedule.Interval.Duration())
			}
		}
	}

	// Temporary variable startRunTime in order to store first execution time
	startRunTime := schedule.StartAt

	if schedule.ExecutionsTimes == nil {
		schedule.ExecutionsTimes = make([]time.Time, schedule.Count)
		for i := 0; i < schedule.Count; i++ {
			schedule.ExecutionsTimes[i] = schedule.StartAt
			schedule.StartAt = schedule.StartAt.Add(schedule.Interval.Duration())
		}
	}

	fmt.Println("\r\nstart:\t\t", startRunTime)
	fmt.Println("end:\t\t", schedule.EndAt)
	fmt.Println("count:\t\t", schedule.Count)
	fmt.Println("interval:\t", schedule.Interval.Duration())
	fmt.Println("executions:\t", schedule.ExecutionsTimes)
	fmt.Println("\r\nTASK SCHEDULE:")
	// calculation times for next executions
	for i := 0; i < schedule.Count; i++ {
		fmt.Printf("TIME #%v [%v]\r\n", i+1, schedule.ExecutionsTimes[i])
	}

	fmt.Println(">>> starting time...", startRunTime)
	// Write into header information about time of first and last task execution
	w.Header().Add("START time", startRunTime.Format(time.RFC1123))
	w.Header().Add("FINISH time", schedule.EndAt.Format(time.RFC1123))

	w.Write([]byte(startRunTime.Format(time.RFC3339)))
}
