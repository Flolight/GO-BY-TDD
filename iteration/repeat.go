package iteration

// Repeat <character> for <repetitions> times
func Repeat(character string, repetitions int) (repeated string) {

	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return
}
