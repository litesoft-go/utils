package uint8s

import (
	"sort"

	// standard libs only above!

	"github.com/litesoft-go/utils/options"
)

var zeroValue []uint8

//noinspection GoUnusedExportedFunction
func ZeroValue() []uint8 {
	return zeroValue
}

//noinspection GoUnusedExportedFunction
func FindIn(in uint8, slice []uint8) (foundIndexOrMinus1 int) {
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
func AddEntry(in uint8, sliceIn []uint8) (sliceOut []uint8, updated bool) {
	sliceOut = sliceIn
	index := FindIn(in, sliceOut)
	if index == -1 {
		sliceOut = append(sliceOut, in)
		updated = true
	}
	return
}

//noinspection GoUnusedExportedFunction
func RemoveEntry(in uint8, sliceIn []uint8) (sliceOut []uint8, updated bool) {
	return RemoveAt(sliceIn, FindIn(in, sliceIn))
}

func RemoveAt(sliceIn []uint8, at int) (sliceOut []uint8, updated bool) {
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
func Copy(src []uint8) []uint8 {
	if src == nil {
		return nil
	}
	dst := make([]uint8, len(src))
	copy(dst, src)
	return dst
}

//noinspection GoUnusedExportedFunction
func SortStable(src []uint8) (out []uint8) {
	out = Copy(src)
	sort.SliceStable(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return
}

//noinspection GoUnusedExportedFunction
func SlicesEqual(in1, in2 []uint8) bool {
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
func AsOptions(src []uint8, ifEmpty string) string {
	collector := options.For(len(src), ifEmpty)
	for _, v := range src {
		collector.Add(ToA(v))
	}
	return collector.Done()
}
