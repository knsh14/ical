package token

import "regexp"

func IsXName(s string) bool {
	isMatch, err := regexp.MatchString(`^X-([\dA-Z]{3}-)?[\dA-Z-]+$`, s)
	if err != nil {
		return false
	}
	return isMatch
}
