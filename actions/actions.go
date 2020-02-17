package actions

import (
	"runtime"
	"strconv"
	"time"
)

func CurrentTime() string {
	return time.Now().String()
}

func CurrentOS() string {
	return runtime.GOOS
}

func CurrentCPU() string {
	return strconv.Itoa(runtime.NumCPU())
}
