package main

import (
	"fmt"
	"sort"
	"strings"

)

func sortString(s string) string {
	freqCount := make(map[rune]int)

	for _, ch := range s {
		freqCount[ch]++
	}

	keys := make([]rune, 0, len(freqCount))
	for ch := range freqCount {
		keys = append(keys, ch)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var result strings.Builder
	for len(result.String()) < len(s) {
		for _, ch := range keys {
			if freqCount[ch] > 0 {
				result.WriteRune(ch)
				freqCount[ch]--
			}
		}
		for i := len(keys) - 1; i >= 0; i-- {
			if freqCount[keys[i]] > 0 {
				result.WriteRune(keys[i])
				freqCount[keys[i]]--
			}
		}
	}

	return result.String()
}

func main() {
	s := "aaaabbbbcccc"
	reordered := sortString(s)
	fmt.Println("Reordered string:", reordered)
}
