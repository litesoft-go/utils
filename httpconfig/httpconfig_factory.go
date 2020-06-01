package httpconfig

import (
	"net"
	"net/http"
	// standard libs only above!
)

//noinspection GoUnusedExportedFunction
func NewClient(timing *Timing) (client *http.Client) {
	if timing != nil {
		return &http.Client{
			Timeout: timing.EndlessRedirectsMaxTime,
			Transport: &http.Transport{
				TLSHandshakeTimeout:   timing.TLSHandshakeTimeout,
				ExpectContinueTimeout: timing.ExpectContinueTimeout,
				ResponseHeaderTimeout: timing.ResponseHeaderTimeout,
				DialContext: (&net.Dialer{
					Timeout:   timing.DialerTimeout,
					KeepAlive: timing.DialerKeepAlive,
				}).DialContext,
			},
		}
	}
	return
}
