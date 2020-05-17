package enums

import (
	"encoding/json"
	"testing"
)

type PrimeEnum struct {
	Enum
	prime bool
}

func (dst *PrimeEnum) UpdateFrom(found IEnum) {
	src := found.(*PrimeEnum)
	dst.prime = src.prime
}

func (dst *PrimeEnum) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(dst, data)
}

var (
	Zero  = PrimeEnum{Enum: New("Zero"), prime: false}
	One   = PrimeEnum{Enum: New("One"), prime: true}
	Two   = PrimeEnum{Enum: New("Two"), prime: true}
	Three = PrimeEnum{Enum: New("Three"), prime: true}
	Four  = PrimeEnum{Enum: New("Four"), prime: false}
)

func init() {
	Add(&Zero, &One, &Two, &Three, &Four)
}

func TestPopulate(t *testing.T) {
	actual := &PrimeEnum{}
	found, err := Populate(actual, "One")
	if !found || (err != nil) {
		t.Errorf("Problem (found=%v): %v\n", found, err)
		return
	}
	if *actual != One {
		t.Errorf("Not Equal:\n"+
			"   expected: %v\n"+
			"     actual: %v\n", One, actual)
	}
}

type Carrier struct {
	Name  string
	Prime *PrimeEnum
}

func TestJson(t *testing.T) {
	c1 := &Carrier{Name: "Fred", Prime: &Two}

	data, err := json.Marshal(c1)
	if err != nil {
		t.Errorf("c1: %w", err)
		return
	}

	c2 := &Carrier{}
	err = json.Unmarshal(data, c2)
	if err != nil {
		t.Errorf("c2: %w", err)
		return
	}
	if (c1.Name != c2.Name) || (c1.Prime == nil) || (c2.Prime == nil) || (*c1.Prime != *c2.Prime) {
		t.Errorf("Not Equal:\n"+
			"  c1: %v\n"+
			"  c2: %v\n", c1, c2)
	}
}
