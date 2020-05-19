package uints

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 uint) uint {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 uint, src2Func Source) uint {
	if pick1 {
		return src1
	}
	return src2Func()
}
