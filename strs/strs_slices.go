package strs

import (
	"sort"

	// standard libs only above!

	"github.com/litesoft-go/utils/options"
)

var zeroValue []string

//noinspection GoUnusedExportedFunction
func ZeroValue() []string {
	return zeroValue
}

//noinspection GoUnusedExportedFunction
func FindIn(in string, slice []string) (foundIndexOrMinus1 int) {
	if len(slice) != 0 {
		for i, nodeID := range slice {
			if nodeID == in {
				return i
			}
		}
	}
	return -1
}

//noinspection GoUnusedExportedFunction
func AddEntry(in string, sliceIn []string) (sliceOut []string, updated bool) {
	sliceOut = sliceIn
	index := FindIn(in, sliceOut)
	if index == -1 {
		sliceOut = append(sliceOut, in)
		updated = true
	}
	return
}

//noinspection GoUnusedExportedFunction
func RemoveEntry(in string, sliceIn []string) (sliceOut []string, updated bool) {
	return RemoveAt(sliceIn, FindIn(in, sliceIn))
}

func RemoveAt(sliceIn []string, at int) (sliceOut []string, updated bool) {
	lastIndex := len(sliceIn) - 1
	if (at < 0) || (lastIndex < at) {
		sliceOut = sliceIn
	} else {
		updated = true
		switch {
		case at == 0:
			sliceOut = sliceIn[1:]
		case at == lastIndex:
			sliceOut = sliceIn[:lastIndex]
		default:
			sliceOut = append(append(sliceOut, sliceIn[:at]...), sliceIn[at+1:]...)
		}
	}
	return
}

//noinspection GoUnusedExportedFunction
func Copy(src []string) []string {
	if src == nil {
		return nil
	}
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

//noinspection GoUnusedExportedFunction
func SortStable(src []string) (out []string) {
	out = Copy(src)
	sort.SliceStable(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return
}

//noinspection GoUnusedExportedFunction
func SlicesEqual(in1, in2 []string) bool {
	l1, l2 := len(in1), len(in2)
	if l1 != l2 {
		return false
	}
	if &in1 != &in2 {
		for i := range in1 {
			if in1[i] != in2[i] {
				return false
			}
		}
	}
	return true
}

//noinspection GoUnusedExportedFunction
func AsOptions(src []string, ifEmpty string) string {
	collector := options.For(len(src), ifEmpty)
	for _, v := range src {
		collector.Add(quote(v))
	}
	return collector.Done()
}
