package uint8s

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 uint8) uint8 {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 uint8, src2Func Source) uint8 {
	if pick1 {
		return src1
	}
	return src2Func()
}
