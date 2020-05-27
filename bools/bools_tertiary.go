package bools

//noinspection GoUnusedExportedFunction
func Tertiary(pick1, src1, src2 bool) bool {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1, src1 bool, src2Func Source) bool {
	if pick1 {
		return src1
	}
	return src2Func()
}
