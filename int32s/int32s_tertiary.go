package int32s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int32) int32 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 int32, src2Func Source) int32 {
	if pick1 {
		return src1
	}
	return src2Func()
}
