package utilities

import "regexp"

func IsCyrillicChar(char string) bool {
	regexpRule := regexp.MustCompile(`\p{Cyrillic}`)
	return regexpRule.MatchString(char)
}

func FindAllIndexes(text string, char rune) []int {
	result := []int{}
	runes := []rune(text)
	for i, textChar := range runes {
		if char == textChar {
			result = append(result, i)
		}
	}
	return result
}

func ReplaceUnderscoreByChar(text []rune, char rune, indexes []int) string {
	result := text
	for _, value := range indexes {
		result[value] = char
	}
	return string(result)
}
