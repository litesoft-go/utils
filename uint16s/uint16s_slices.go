package uint16s

import (
	"sort"

	// standard libs only above!

	"github.com/litesoft-go/utils/options"
)

var zeroValue []uint16

//noinspection GoUnusedExportedFunction
func ZeroValue() []uint16 {
	return zeroValue
}

//noinspection GoUnusedExportedFunction
func FindIn(in uint16, slice []uint16) (foundIndexOrMinus1 int) {
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
func AddEntry(in uint16, sliceIn []uint16) (sliceOut []uint16, updated bool) {
	sliceOut = sliceIn
	index := FindIn(in, sliceOut)
	if index == -1 {
		sliceOut = append(sliceOut, in)
		updated = true
	}
	return
}

//noinspection GoUnusedExportedFunction
func RemoveEntry(in uint16, sliceIn []uint16) (sliceOut []uint16, updated bool) {
	return RemoveAt(sliceIn, FindIn(in, sliceIn))
}

func RemoveAt(sliceIn []uint16, at int) (sliceOut []uint16, updated bool) {
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
func Copy(src []uint16) []uint16 {
	if src == nil {
		return nil
	}
	dst := make([]uint16, len(src))
	copy(dst, src)
	return dst
}

//noinspection GoUnusedExportedFunction
func SortStable(src []uint16) (out []uint16) {
	out = Copy(src)
	sort.SliceStable(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return
}

//noinspection GoUnusedExportedFunction
func SlicesEqual(in1, in2 []uint16) bool {
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
func AsOptions(src []uint16, ifEmpty string) string {
	collector := options.For(len(src), ifEmpty)
	for _, v := range src {
		collector.Add(ToA(v))
	}
	return collector.Done()
}
