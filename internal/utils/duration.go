package utils

import "time"

func RoundDuration(d time.Duration) time.Duration {
	if d < 1000*time.Nanosecond {
		return d.Round(time.Nanosecond)
	} else if d < 1000*time.Microsecond {
		return d.Round(time.Nanosecond)
	} else if d < 1000*time.Millisecond {
		return d.Round(time.Microsecond)
	} else if d < 1000*time.Second {
		return d.Round(time.Millisecond)
	} else {
		return d.Round(time.Second)
	}
}
