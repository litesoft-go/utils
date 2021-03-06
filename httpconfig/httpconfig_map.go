package httpconfig

import (
	"net/http"
	"sync"
	// standard libs only above!
)

var timingsByTimeOut = struct {
	sync.RWMutex
	m map[int]*Timing
}{m: make(map[int]*Timing)}

func timingByTimeout(timeoutSeconds int) *Timing {
	timingsByTimeOut.RLock()
	timing := timingsByTimeOut.m[timeoutSeconds]
	timingsByTimeOut.RUnlock()
	if timing == nil {
		timing = NewTiming(timeoutSeconds)
		timingsByTimeOut.Lock()
		timingsByTimeOut.m[timeoutSeconds] = timing
		timingsByTimeOut.Unlock()
	}
	return timing
}

var timingsByName = struct {
	sync.RWMutex
	m map[string]*Timing
}{m: make(map[string]*Timing)}

func SetNamedTiming(name string, timing *Timing) *Timing {
	timingsByName.Lock()
	if timing == nil {
		delete(timingsByName.m, name)
	} else {
		timingsByName.m[name] = timing
	}
	timingsByName.Unlock()
	return timing
}

func SetNamedTimeout(name string, timeoutSeconds int) *Timing {
	return SetNamedTiming(name, timingByTimeout(timeoutSeconds))
}

func getNamed(name string) *Timing {
	timingsByName.RLock()
	timing := timingsByName.m[name]
	timingsByName.RUnlock()
	return timing
}

func GetNamedTimeout(name string, defaultTimeoutSeconds int) *Timing {
	timing := getNamed(name)
	if timing == nil {
		timing = SetNamedTimeout(name, defaultTimeoutSeconds)
	}
	return timing
}

func GetNamedTiming(name string, defaultTiming *Timing) (timing *Timing) {
	timing = getNamed(name)
	if (timing == nil) && (defaultTiming != nil) {
		timing = SetNamedTiming(name, defaultTiming)
	}
	return
}

//noinspection GoUnusedExportedFunction
func ClientByNameWithDefaultTiming(name string, defaultTiming *Timing) *http.Client {
	return ClientByTiming(GetNamedTiming(name, defaultTiming))
}

//noinspection GoUnusedExportedFunction
func ClientByNameWithDefaultTimeout(name string, defaultTimeoutSeconds int) *http.Client {
	return ClientByTiming(GetNamedTimeout(name, defaultTimeoutSeconds))
}

//noinspection GoUnusedExportedFunction
func ClientByTimeout(timeoutSeconds int) *http.Client {
	return ClientByTiming(timingByTimeout(timeoutSeconds))
}

var clientsByTiming = struct {
	sync.RWMutex
	m map[Timing]*http.Client
}{m: make(map[Timing]*http.Client)}

func ClientByTiming(timing *Timing) (client *http.Client) {
	if timing != nil {
		clientsByTiming.RLock()
		client = clientsByTiming.m[*timing]
		clientsByTiming.RUnlock()
		if client == nil {
			client = NewClient(timing)
			clientsByTiming.Lock()
			clientsByTiming.m[*timing] = client
			clientsByTiming.Unlock()
		}
	}
	return
}
