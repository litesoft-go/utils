package uint32s

//noinspection GoUnusedExportedFunction
func FindIn(in uint32, slice []uint32) (foundIndexOrMinus1 int) {
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
func AddEntry(in uint32, sliceIn []uint32) (sliceOut []uint32, updated bool) {
	sliceOut = sliceIn
	index := FindIn(in, sliceOut)
	if index == -1 {
		sliceOut = append(sliceOut, in)
		updated = true
	}
	return
}

//noinspection GoUnusedExportedFunction
func RemoveEntry(in uint32, sliceIn []uint32) (sliceOut []uint32, updated bool) {
	return RemoveAt(sliceIn, FindIn(in, sliceIn))
}

func RemoveAt(sliceIn []uint32, at int) (sliceOut []uint32, updated bool) {
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