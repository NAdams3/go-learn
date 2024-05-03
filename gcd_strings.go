package main

import "strings"

func gcdOfStrings(str1, str2 string) string {
	// take smallest string remove one letter until other contains substring or first string is ""

	s1Len := len(str1)
	s3Len := len(str2)
	shortest := str1
	longest := str2

	if s3Len < s1Len {
		shortest = str2
		longest = str1
	}

	for i := range len(shortest) {
		if strings.Contains(longest, shortest[i:]) {
			return shortest[i:]
		}
	}
	return ""
}
