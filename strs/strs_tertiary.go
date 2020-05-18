package strs

// DO NOT USE WITH ERRORS!
//
// strs.Tertiary(err != nil, err.Error(), "N/A") leads to nil pointer dereference in case of nil error
// as soon as second parameter evaluated in any case
//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 string) string {
	if pick1 {
		return src1
	}
	return src2
}

//noinspection GoUnusedExportedFunction
func TertiaryFunc(pick1 bool, src1 string, src2Func Source) string {
	if pick1 {
		return src1
	}
	return src2Func()
}
