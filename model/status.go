package model

type Status string

const (
	Running Status = "running"
	Failed  Status = "failed"
	Succeed Status = "succeed"
)
