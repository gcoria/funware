package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	words := []string{"cat", "bat", "ads", "sad", "dog"}
	fmt.Println(GroupAnagrams(words))
}

func GroupAnagrams(words []string) map[string][]string {
	grouped := make(map[string][]string)
	for _, word := range words {
		chars := strings.Split(word, "")
		slices.Sort(chars)
		key := strings.Join(chars, "")
		grouped[key] = append(grouped[key], word)
	}
	return grouped
}
