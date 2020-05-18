package ints

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int) int {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 int, src2Func Source) int {
	if pick1 {
		return src1
	}
	return src2Func()
}
