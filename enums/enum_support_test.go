package enums

import (
	"encoding/json"
	"fmt"
	"testing"
	// standard libs only above!
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

	wilma := &Carrier{Name: "Wilma", Prime: &One}

	err := roundTrip(wilma)

	if err != nil {
		t.Errorf("Wilma: %w", err)
		return
	}

	fred := &Carrier{Name: "Fred"}

	err = roundTrip(fred)

	if err != nil {
		t.Errorf("Fred: %w", err)
		return
	}
}

func roundTrip(in *Carrier) error {
	data, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("in: %w", err)
	}

	out := &Carrier{}
	err = json.Unmarshal(data, out)
	if err != nil {
		return fmt.Errorf("out: %w", err)
	}

	if (in.Name != out.Name) || !AreSame(in.Prime, out.Prime) {
		err = fmt.Errorf("Not Equal:\n"+
			"   in: %v\n"+
			"  out: %v\n", in, out)
	}

	return err
}
