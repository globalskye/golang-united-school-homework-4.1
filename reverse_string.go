package reverse_string

func ReverseString(input string) (output string) {
	//result := make([]string, len(input))
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return input
	}
	i := len(input) + 1
	for {
		if i != 1 {
			//	i--
			output = output + input[i-2:i-1]
			i--
		} else {
			break
		}

	}
	return output
}
