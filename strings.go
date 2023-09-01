package commons

func IsEmpty(val ...string) bool {
	for _, v := range val {
		if len(v) == 0 {
			return true
		}
	}
	return false
}
