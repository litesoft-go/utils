package httpconfig

import (
	"strconv"
	"time"
	// standard libs only above!
)

type Timing struct {
	fromTimeoutSeconds      int
	DialerTimeout           time.Duration
	DialerKeepAlive         time.Duration
	TLSHandshakeTimeout     time.Duration
	EndlessRedirectsMaxTime time.Duration
	ExpectContinueTimeout   time.Duration
	ResponseHeaderTimeout   time.Duration
}

//noinspection GoUnusedExportedFunction
func NewTiming(timeoutSeconds int) *Timing {
	if timeoutSeconds < 1 {
		panic("Attempt to create a New httpconfig.Timing with timeouts of: " + strconv.Itoa(timeoutSeconds))
	}

	return &Timing{
		fromTimeoutSeconds:      timeoutSeconds,
		DialerTimeout:           duration(timeoutSeconds),
		DialerKeepAlive:         duration(timeoutSeconds),
		TLSHandshakeTimeout:     duration(timeoutSeconds),
		EndlessRedirectsMaxTime: duration(2 * timeoutSeconds),
		ExpectContinueTimeout:   fractionalDuration(0.4, timeoutSeconds),
		ResponseHeaderTimeout:   fractionalDuration(0.3, timeoutSeconds),
	}
}

//noinspection GoUnusedExportedFunction
func (in *Timing) IsBasedOn(timeoutSeconds int) bool {
	if in != nil {
		return timeoutSeconds == in.fromTimeoutSeconds
	}
	return false
}

func duration(timeoutSeconds int) time.Duration {
	return time.Duration(timeoutSeconds) * time.Second
}

func fractionalDuration(fraction float32, timeoutSeconds int) time.Duration {
	float := (fraction * float32(timeoutSeconds)) + 0.9
	return duration(int(float))
}
