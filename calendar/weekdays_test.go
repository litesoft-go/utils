package calendar

import (
	"encoding/json"
	"fmt"
	"testing"

	// standard libs only above!

	"github.com/litesoft-go/utils/enums"
)

func TestRegistered(t *testing.T) {
	expected := []enums.IEnum{&Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday, &Sunday}
	actual := enums.GetRegisteredMembers(&Monday)

	if len(expected) != len(actual) {
		t.Errorf("Not Equal lengths:\n"+
			"   expected: %v\n"+
			"     actual: %v\n", expected, actual)
		return
	}
	for i := range expected {
		e, a := expected[i], actual[i]
		if e != a {
			t.Errorf("Not Equal %s:\n"+
				"   expected: %v\n"+
				"     actual: %v\n", e.Name(), e, a)
			return
		}
	}
}

type Carrier struct {
	Name       string   `json:"Name,omitempty"`
	StartDay   Weekday  `json:"startDay"`
	EndDay     *Weekday `json:"endDay,omitempty"`
	FourDayOff *Weekday `json:"fourDayOff,omitempty"`
}

func TestJson(t *testing.T) {
	c1 := &Carrier{Name: "Fred", StartDay: Monday, FourDayOff: &Wednesday}

	d1, err := json.Marshal(c1)
	if err != nil {
		t.Errorf("c1 -> JSON: %w", err)
		return
	}

	c2 := &Carrier{}
	err = json.Unmarshal(d1, c2)
	if err != nil {
		t.Errorf("JSON -> c2: %w", err)
		return
	}

	d2, err := json.Marshal(c2)
	if err != nil {
		t.Errorf("c2 -> JSON: %w", err)
		return
	}

	s1, s2 := string(d1), string(d2)
	if s1 != s2 { // (c1.Name != c2.Name) || (c1.Prime == nil) || (c2.Prime == nil) || (*c1.Prime != *c2.Prime) {
		t.Errorf("Not Equal:\n"+
			"  c1: %v\n"+
			"  c2: %v\n", s1, s2)
		return
	}
	fmt.Println(s1)
}
