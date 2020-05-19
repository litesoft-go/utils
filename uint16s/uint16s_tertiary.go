package uint16s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 uint16) uint16 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 uint16, src2Func Source) uint16 {
	if pick1 {
		return src1
	}
	return src2Func()
}
