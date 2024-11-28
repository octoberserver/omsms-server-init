package main

import "regexp"

func isURL(str string) bool {
	re, err := regexp.Compile(`https?://[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)`)
	if err != nil {
		panic("Failed to compile regex " + err.Error())
	}
	return re.MatchString(str)
}

func checkStringMatches(str string, options []string) bool {
	for _, option := range options {
		if str == option {
			return true
		}
	}
	return false
}
