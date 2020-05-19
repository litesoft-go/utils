package uint64s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 uint64) uint64 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 uint64, src2Func Source) uint64 {
	if pick1 {
		return src1
	}
	return src2Func()
}
