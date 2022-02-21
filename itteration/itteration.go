package itteration

func Repeat(word string, count int) string {
	result := word
	for i := 1; i < count; i++ {
		result += " " + word
	}
	return result
}
