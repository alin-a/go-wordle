package game

import "strings"

const (
	IsInCorrectPosition int = 2
	OccursInWord            = 1
	DoesNotOccurInWord      = 0
)

func CompareInputAndWord(input string, word string) [5]int {
	return matchOccurrence(toMapOfCharacterAndIndices(strings.ToUpper(input)), toMapOfCharacterAndIndices(strings.ToUpper(word)))
}

func matchOccurrence(inputMap map[string][]int, wordMap map[string][]int) [5]int {
	var result [5]int

	for character, indices := range inputMap {
		indicesInWord := wordMap[character]
		frequencyInWord := len(indicesInWord)

		frequencyInInput := len(indices)
		if frequencyInWord == 0 { // character is not in word
			for _, index := range indices {
				result[index] = DoesNotOccurInWord
			}
		} else if frequencyInWord < frequencyInInput { // character is in word with lower frequency than in input
			counter := 0
			// first: mark all characters in correct position
			for _, index := range indicesInWord {
				if contains(indices, index) {
					result[index] = IsInCorrectPosition
					counter++
				}
			}
			// second: if the input contains more of the character, mark their indices (from left to right) until
			// the frequency in the word is reached
			for _, index := range indices {
				if counter < frequencyInWord {
					result[index] = OccursInWord
					counter++
				}
			}
		} else { // character is in word with same or higher frequency than in input
			for _, index := range indices {
				if contains(indicesInWord, index) {
					result[index] = IsInCorrectPosition
				} else {
					result[index] = OccursInWord
				}
			}
		}
	}
	return result
}

func contains(elems []int, value int) bool {
	for _, elem := range elems {
		if value == elem {
			return true
		}
	}
	return false
}

func toMapOfCharacterAndIndices(word string) map[string][]int {
	result := make(map[string][]int)
	for index, character := range word {
		charStr := string(character)
		result[charStr] = append(result[charStr], index)
	}
	return result
}
