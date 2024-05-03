package main

import "strings"

func compressChars(chars []string) int {
	// have a concate string of unique characters and a count array of repition counts for each char in concat string
	concat := ""
	counts := make([]int, 0)
	for i := range len(chars) {
		if !strings.Contains(concat, string(chars[i])) {
			concat += string(chars[i])
			counts = append(counts, 1)
		} else {
			ind := strings.Index(concat, string(chars[i]))
			counts[ind]++
		}
	}

	length := 0
	for _, count := range counts {
		if count > 1 {
			length++
		}
	}

	return length + len(concat)
}
