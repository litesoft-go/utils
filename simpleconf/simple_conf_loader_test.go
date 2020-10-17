package simpleconf

import (
	"reflect"
	"strings"
	"testing"
)

type MapOf struct {
	keyValues map[string][]string
}

func (in *MapOf) mapOf(key string, values ...string) *MapOf {
	if (key != "") && (len(values) != 0) {
		in.keyValues[key] = values
	}
	return in
}

func mapOf(key string, values ...string) *MapOf {
	mo := MapOf{keyValues: make(map[string][]string)}
	return mo.mapOf(key, values...)
}

func expectedError(t *testing.T, expectedErr string, lines ...string) {
	contents := strings.Join(lines, "\n") + "\n"
	config, err := LoadReader(strings.NewReader(contents))
	if err == nil {
		t.Errorf("Expecting err, but got contents:\n%s", config)
		return
	}
	actualErr := err.Error()
	if actualErr != expectedErr {
		t.Errorf("Not Equal:\n"+
			"   expected: '%s'\n"+
			"     actual: '%s'\n", expectedErr, actualErr)
	}
}

func check(t *testing.T, mapOf *MapOf, lines ...string) {
	contents := strings.Join(lines, "\n") + "\n"
	config, err := LoadReader(strings.NewReader(contents))
	if err != nil {
		t.Errorf("Not expecting err, but got error: %s\n"+
			"on contents:\n%s", err, contents)
		return
	}
	expectedKeyValues := mapOf.keyValues
	if !reflect.DeepEqual(expectedKeyValues, config.keyValues) {
		t.Errorf("Not Equal:\n"+
			"   expected: '%s'\n"+
			"     actual: '%s'\n", expectedKeyValues, config.keyValues)
	}
}

func TestEmpty(t *testing.T) {
	check(t, mapOf(""), "")
	check(t, mapOf(""),
		"",
		"! comment",
		"# comment",
		"!",
		"#",
		"  !",
		"  #",
		"")
}

func TestOneEntry(t *testing.T) {
	check(t, mapOf("xy", "zzy"), "xy:zzy")
}

func TestThreeEntry(t *testing.T) {
	check(t,
		mapOf("xy", "zzy").
			mapOf("abc", "def", "ghi").
			mapOf("fred", "wilma"),
		"",
		"# xy:",
		"xy:zzy",
		"",
		"# abc:",
		"abc:",
		" def",
		" # 2nd",
		" ghi",
		"",
		"# F & W:",
		"fred:",
		"  # spouse:",
		"     wilma",
	)
}
func TestErrors(t *testing.T) {
	expectedError(t, "list 'abc' indicated, but no entries added",
		"abc:",
		" # 2nd",
	)
	expectedError(t, "no list indicated, but indented line encountered:  abc",
		" abc",
	)
	expectedError(t, "no separator ':' or '=' found in line: abc",
		"abc",
	)
	expectedError(t, "'key' empty in line: =abc",
		"=abc",
	)
}
