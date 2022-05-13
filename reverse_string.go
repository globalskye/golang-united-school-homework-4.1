package reverse_string

func ReverseString(input string) (output string) {
	//result := make([]string, len(input))
	i := len(input)

	for {
		if i != 1 {
			i--
			output = output + input[i-1:i]
		} else {
			break
		}

	}
	return output
}
