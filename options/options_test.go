package options

import (
	"testing"
)

func check(t *testing.T, ifEmpty, expected string, adds ...string) {
	collector := For(len(adds), ifEmpty)
	for _, v := range adds {
		collector.Add(v)
	}
	actual := collector.Done()
	if actual != expected {
		t.Errorf("Not Equal:\n"+
			"   expected: '%s'\n"+
			"     actual: '%s'\n", expected, actual)
	}
}

func TestEmpty(t *testing.T) {
	check(t, "", "")
	check(t, "N/A", "N/A")
}

func TestOneEntry(t *testing.T) {
	expected := "xyzzy"
	check(t, "", expected, expected)
	check(t, "N/A", expected, expected)
}

func TestTwoEntry(t *testing.T) {
	expected, adds := "xy or zzy", []string{"xy", "zzy"}
	check(t, "", expected, adds...)
	check(t, "N/A", expected, adds...)
}

func TestThreeEntry(t *testing.T) {
	expected, adds := "xy, zz, or y", []string{"xy", "zz", "y"}
	check(t, "", expected, adds...)
	check(t, "N/A", expected, adds...)
}

func TestIncomplete(t *testing.T) {
	expected := MISSING_VALUE
	actual := For(1, "").Done()
	if actual != expected {
		t.Errorf("Not Equal:\n"+
			"   expected: '%s'\n"+
			"     actual: '%s'\n", expected, actual)
	}
}
