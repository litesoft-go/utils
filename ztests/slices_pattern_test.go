package ztests

import (
	"testing"

	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
)

func TestSlices(t *testing.T) {
	var sliceIn []int

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{1, 2, 3, 5, 8}, false, "add existing").
		check(ints.AddEntry, 3)

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{1, 2, 3, 5, 8, 13}, true, "add new").
		check(ints.AddEntry, 13)

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{1, 2, 3, 5, 8}, false, "remove non-existing").
		check(ints.RemoveEntry, 13)

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{2, 3, 5, 8}, true, "remove 1st").
		check(ints.RemoveEntry, 1)

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{1, 2, 3, 5}, true, "remove last").
		check(ints.RemoveEntry, 8)

	sliceIn = Copy(1, 2, 3, 5, 8)
	create(t, sliceIn, []int{1, 3, 5, 8}, true, "remove middle").
		check(ints.RemoveEntry, 2)
}

type answers struct {
	t                       *testing.T
	what                    string
	orig, sliceIn, expected []int
	expectedUpdated         bool
}

func create(t *testing.T, sliceIn, expected []int, expectedUpdated bool, what string) *answers {
	orig := Copy(sliceIn...)
	return &answers{t: t, what: what, orig: orig, sliceIn: sliceIn,
		expected: expected, expectedUpdated: expectedUpdated}
}

func (a *answers) check(method func(int, []int) ([]int, bool), toFind int) {
	sliceOut, updated := method(toFind, a.sliceIn)
	if !Equal(sliceOut, a.expected) {
		a.t.Errorf("%s\n  Not Equal result slices:\n"+
			"   expected: %v\n"+
			"     actual: %v\n", a.what, a.expected, sliceOut)
		return
	}
	if updated != a.expectedUpdated {
		a.t.Errorf("%s\n  Not Equal updated:\n"+
			"   expected: %v\n"+
			"     actual: %v\n", a.what, a.expectedUpdated, updated)
		return
	}
	if Equal(a.orig, a.sliceIn) {
		return
	}
	if updated {
		a.t.Errorf("%s\n  Not Equal (updated) orig & sliceIn:\n"+
			"   sliceIn: %v\n"+
			"      orig: %v\n", a.what, a.sliceIn, a.orig)
	} else {
		a.t.Errorf("%s\n  Not Equal (!updated) orig & sliceIn:\n"+
			"   sliceIn: %v\n"+
			"      orig: %v\n", a.what, a.sliceIn, a.orig)
	}
}

func Copy(in ...int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
