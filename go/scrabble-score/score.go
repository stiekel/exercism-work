package scrabble

import "strings"

func Score(s string) int {
	value := 0
	valueMap := map[int]string{
		1: "AEIOULNRST", 2: "DG", 3: "BCMP", 4: "FHVWY", 5: "K", 8: "JX", 10: "QZ"}
	for i := range s {
		for j := range valueMap {
			index := strings.IndexAny(valueMap[j], strings.ToUpper(s[i:i+1]))
			if index != -1 {
				value += j
			}
		}
	}
	return value
}
