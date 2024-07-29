package iteration

func Repeat(toRepeat string, repetions int) string {
	repeatedString := toRepeat
	for range repetions-1 {
		repeatedString += toRepeat
	}
	return repeatedString
}

func Repeat2(toRepeat string, repetions int) string {
	var repeated string
	for i := 0; i < repetions + 1; i++ {
repeated += toRepeat
	}
	return repeated
}
