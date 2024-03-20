package iteration

const charCount int = 5

func Repeat(character string) string {
	var repeated string

	for i := 0; i < charCount; i++ {
		repeated += character
	}

	return repeated
}
