func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	index := 0
	for _, tt := range t {
		if byte(tt) == s[index] {
			if index == len(s)-1 {
				return true
			}
			index += 1
		}
	}
	return false
}