package int16s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int16) int16 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 int16, src2Func Source) int16 {
	if pick1 {
		return src1
	}
	return src2Func()
}
