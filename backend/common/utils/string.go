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
