package iteration

func Repeat(character string, numTimesCharRepeats int) string {
	var repeated string

	for i := 0; i < numTimesCharRepeats; i++ {
		repeated += character
	}

	return repeated
}
