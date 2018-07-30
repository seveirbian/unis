package cmd

// generate a substring
func Substring(str string, start int, end int) string {
	stringSlice := []rune(str)
	stringLen := len(stringSlice)

	if start < 0 || start >= stringLen {
		return ""
	} else if end < 0 || end >= stringLen || end < start {
		return ""
	} else {
		return string(stringSlice[start:end])
	}
}

// generate empty string
func EmptyString(length int) string {
	var stringSlice []rune

	for i := 0; i < length; i++ {
		stringSlice = append(stringSlice, rune(' '))
	}

	return string(stringSlice[:])
}
