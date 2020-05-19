package int8s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int8) int8 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 int8, src2Func Source) int8 {
	if pick1 {
		return src1
	}
	return src2Func()
}
