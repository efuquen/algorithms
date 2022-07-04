package stopwatch

import "time"

type Stopwatch struct {
	start int64
}

func NewStopwatch() Stopwatch {
	return Stopwatch{
		start: time.Now().UnixMilli(),
	}
}

func (s *Stopwatch) ElapsedTime() float64 {
	now := time.Now().UnixMilli()
	return float64(now-s.start) / 1000.0
}
