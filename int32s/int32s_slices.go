package int32s

//noinspection GoUnusedExportedFunction
func FindIn(in int32, slice []int32) (foundIndexOrMinus1 int) {
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
func AddEntry(in int32, sliceIn []int32) (sliceOut []int32, updated bool) {
	sliceOut = sliceIn
	index := FindIn(in, sliceOut)
	if index == -1 {
		sliceOut = append(sliceOut, in)
		updated = true
	}
	return
}

//noinspection GoUnusedExportedFunction
func RemoveEntry(in int32, sliceIn []int32) (sliceOut []int32, updated bool) {
	return RemoveAt(sliceIn, FindIn(in, sliceIn))
}

func RemoveAt(sliceIn []int32, at int) (sliceOut []int32, updated bool) {
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
