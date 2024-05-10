package piscine

import (
	"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	words := strings.Split(text, " ")

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	wordCounts := make([]struct {
		word  string
		count int
	}, 0, len(wordCount))
	for word, count := range wordCount {
		wordCounts = append(wordCounts, struct {
			word  string
			count int
		}{word, count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].count != wordCounts[j].count {
			return wordCounts[i].count > wordCounts[j].count
		}
		return wordCounts[i].word != "" && wordCounts[i].word < wordCounts[j].word
	})

	result := make(map[string]int, n)
	for i := 0; i < n && i < len(wordCounts); i++ {
		result[wordCounts[i].word] = wordCounts[i].count
	}
	return result
}
