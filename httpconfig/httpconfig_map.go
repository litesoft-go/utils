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

//noinspection GoUnusedExportedFunction
func ClientByTimeout(timeoutSeconds int) *http.Client {
	timingsByTimeOut.RLock()
	timing := timingsByTimeOut.m[timeoutSeconds]
	timingsByTimeOut.RUnlock()
	if timing == nil {
		timing = NewTiming(timeoutSeconds)
		timingsByTimeOut.Lock()
		timingsByTimeOut.m[timeoutSeconds] = timing
		timingsByTimeOut.Unlock()
	}
	return ClientByTiming(timing)
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
