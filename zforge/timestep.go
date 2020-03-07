package zforge

import (
	"time"
)

type TimeStep struct {
	step time.Duration
}

func NewTimeStep(dur time.Duration) TimeStep {
	return TimeStep{step: dur}
}

func (ts TimeStep) GetSeconds() float64 {
	return ts.step.Seconds()
}

func (ts TimeStep) GetMilliseconds() int64 {
	return ts.step.Milliseconds()
}
