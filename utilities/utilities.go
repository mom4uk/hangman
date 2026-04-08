package utilities

import "regexp"

func IsCyrillicChar(char string) bool {
	regexpRule := regexp.MustCompile(`\p{Cyrillic}`)
	return regexpRule.MatchString(char)
}
