package maps

func NewStrings(initialCapacity int) map[string]string {
	return make(map[string]string, initialCapacity)
}

func NonNilStrings(pMap map[string]string) map[string]string {
	if pMap == nil {
		pMap = NewStrings(1)
	}
	return pMap
}

//noinspection GoUnusedExportedFunction
func CopyStrings(src, dst map[string]string) map[string]string {
	return copyStrings(src, NonNilStrings(dst))
}

//noinspection GoUnusedExportedFunction
func NilableCopyStrings(src map[string]string) map[string]string {
	if src == nil {
		return nil
	}
	return copyStrings(src, NewStrings(len(src)))
}

func copyStrings(src, dst map[string]string) map[string]string {
	for key, value := range src {
		dst[key] = value
	}
	return dst
}
