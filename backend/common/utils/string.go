package utils

import "regexp"

func validateRegex(str string, regex string) bool {
	regexp, err := regexp.Compile(regex)
	if err != nil {
		return false
	}
	return !regexp.Match([]byte(str))
}

// CheckValidUsernameCharacters check to see if the given string has special characters. That is, if it
//	will return false if it contains anything other than letters, numbers, or underscores.
func CheckValidUsernameCharacters(str string) bool {
	return validateRegex(str, `^[A-Za-z0-9_]+$`)
}

// CheckValidPasswordCharacters will check to see if the given password only consists of the allowed
// characters.
func CheckValidPasswordCharacters(str string) bool {
	return validateRegex(str, `^[A-Za-z0-9\[\*\.\!\@\#\$\%\^\&\(\)\{\}\[\]\:\;\<\>\,\.\?\/\~\_\+\-\=\|\\\]]*$`)
}

// FindFirstDifference will find the first difference between two strings. It will use str1 as the base
// and then find the first instance that str2 is different from str1. In the case where there is no difference,
// it will return the length of str1
func FindFirstDifference(str1, str2 string) int {
	idx := 0
	for ; idx < len(str1) && idx < len(str2); idx++ {
		if str1[idx] != str2[idx] {
			return idx
		}
	}
	return idx
}
