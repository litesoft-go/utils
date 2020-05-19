package int64s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int64) int64 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 int64, src2Func Source) int64 {
	if pick1 {
		return src1
	}
	return src2Func()
}
