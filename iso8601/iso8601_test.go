package iso8601

import (
	"testing"
)

func TestEmptyStringParse(t *testing.T) {
	zTime, err := ZmillisFromString("")
	if err == nil {
		t.Errorf("Expected Error. but got none, time was: %v", zTime)
	}
}

const (
	DateTime1 = "2011-01-16T14:03:02.001Z"
	DateTime2 = "2019-07-30T08:05:01.000Z"
	DateTime3 = "2000-02-29T07:06:05.000Z"
)

func TestRoundTrip(t *testing.T) {
	checkRT(t, DateTime1)
	checkRT(t, DateTime2)
	checkRT(t, DateTime3)
}

func checkRT(t *testing.T, expected string) {
	zTime, err := ZmillisFromString(expected)
	if err != nil {
		t.Errorf("Unexpected Error with '%s': %v", expected, err)
	}
	actual := ZmillisToString(&zTime)
	if actual != expected {
		t.Errorf("Roundtrip '%s' != '%s' (orig)", expected, actual)
	}
}
