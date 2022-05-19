package reverse_string

func ReverseString(input string) (output string) {
	//result := make([]string, len(input))
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return input
	}
	result := []rune(input)

	i := len([]rune(input)) + 1
	for {
		if i != 1 {
			//	i--
			output += string(result[i-2 : i-1])
			i--
		} else {
			break
		}

	}
	return output
}
