package iteration

func Repeat(character string, repeatedCharCount int) string {
	var repeated string

	for i := 0; i < repeatedCharCount; i++ {
		repeated += character
	}

	return repeated
}
