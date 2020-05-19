package uint32s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 uint32) uint32 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 uint32, src2Func Source) uint32 {
	if pick1 {
		return src1
	}
	return src2Func()
}
