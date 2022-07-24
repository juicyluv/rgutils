package ptr

import "time"

func String(val string) *string {
	return &val
}

func Float32(val float32) *float32 {
	return &val
}

func Float64(val float64) *float64 {
	return &val
}

func Bool(val bool) *bool {
	return &val
}

func Time(val time.Time) *time.Time {
	return &val
}
