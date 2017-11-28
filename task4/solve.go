package main

import (
	"unicode"
	)

func RemoveEven(input [] int) [] int {
	result := make([] int, 0) // create an empty slice with non-zero length, use the builtin make.

	for i := 0; i < len(input); i++ {
		if input[i] % 2 == 1 {
			result = append(result, input[i])
		}
	}
	return result
}

func PowerGenerator(number int) (func() int) {
	pow := 1
	return func() int {
		pow =  number * pow
		return pow
	}
}

func DifferentWordsCount (input string) int {
	allWords := make([]string, 1) // create an empty slice length 1, use the builtin make.
	result := make(map[string] int)

	var (
		isWord = true
		pos = 0
	)

	for i := 0; i < len(input); i++ {

		if unicode.IsLetter(rune(input[i])) {
			isWord = true
			allWords[ pos ] = allWords[ pos ] + string(unicode.ToLower(rune(input[i])))

		} else if isWord {
			isWord = false
			allWords = append(allWords, "")
			pos ++
		}
	}

	for i := range allWords {
		if allWords[i] != "" {
			result[allWords[i]]++
		}
	}

	return len(result)
}
/*
func main() {

	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]


	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8


	fmt.Println(DifferentWordsCount("HI hi good day"))
	// Должно напечатать 2

}*/