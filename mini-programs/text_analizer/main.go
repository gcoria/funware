package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type TextStats struct {
	CharCount         int
	WordCount         int
	SentenceCount     int
	ParagraphCount    int
	AverageWordLength float64
	CommonWords       map[string]int
	WordLengths       map[int]int
	NumbersFound      []int
	ReadabilityScore  float64
	UniqueWords       int
	LetterFrequency   map[rune]int
}

func Analyze(text string) *TextStats {
	stats := &TextStats{
		CommonWords:     make(map[string]int),
		WordLengths:     make(map[int]int),
		LetterFrequency: make(map[rune]int),
	}

	stats.CharCount = len(text)

	paragraphs := regexp.MustCompile(`\n\s*\n`).Split(text, -1)
	stats.ParagraphCount = len(paragraphs)

	sentenceRegex := regexp.MustCompile(`[.!?]+`)
	stats.SentenceCount = len(sentenceRegex.FindAllString(text, -1))

	wordRegex := regexp.MustCompile(`[\p{L}\p{N}]+`)
	words := wordRegex.FindAllString(text, -1)
	stats.WordCount = len(words)

	totalWordLength := 0
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		stats.CommonWords[lowerWord]++
		wordLen := len(lowerWord)
		stats.WordLengths[wordLen]++
		totalWordLength += wordLen
	}

	stats.UniqueWords = len(stats.CommonWords)

	if stats.WordCount > 0 {
		stats.AverageWordLength = float64(totalWordLength) / float64(stats.WordCount)
	}

	for _, r := range text {
		if unicode.IsLetter(r) {
			stats.LetterFrequency[unicode.ToLower(r)]++
		}
	}

	stats.NumbersFound = extractNumbers(text)

	stats.ReadabilityScore = calculateReadability(text, stats)

	return stats
}

func extractNumbers(text string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text, -1)

	numbers := []int{}
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func calculateReadability(text string, stats *TextStats) float64 {
	if stats.WordCount == 0 || stats.SentenceCount == 0 {
		return 0
	}

	wordsPerSentence := float64(stats.WordCount) / float64(stats.SentenceCount)
	score := 206.835 - (1.015 * wordsPerSentence) - (84.6 * stats.AverageWordLength / 5)
	return score
}

func (stats *TextStats) DisplayStats() {
	fmt.Println("=== Text Statistics ===")
	fmt.Printf("Character count: %d\n", stats.CharCount)
	fmt.Printf("Word count: %d\n", stats.WordCount)
	fmt.Printf("Unique words: %d\n", stats.UniqueWords)
	fmt.Printf("Sentence count: %d\n", stats.SentenceCount)
	fmt.Printf("Paragraph count: %d\n", stats.ParagraphCount)
	fmt.Printf("Average word length: %.2f\n", stats.AverageWordLength)
	fmt.Printf("Readability score: %.2f\n", stats.ReadabilityScore)

	fmt.Println("\n=== Top 5 Most Common Words ===")
	topWords := getTopItems(stats.CommonWords, 5)
	for _, item := range topWords {
		fmt.Printf("%s: %d occurrences\n", item.Key, item.Value)
	}

	fmt.Println("\n=== Top 5 Most Common Letters ===")
	topLetters := getTopItemsRune(stats.LetterFrequency, 5)
	for _, item := range topLetters {
		fmt.Printf("%c: %d occurrences\n", item.Key, item.Value)
	}

	if len(stats.NumbersFound) > 0 {
		fmt.Println("\n=== Numbers Found ===")
		fmt.Printf("Count: %d\n", len(stats.NumbersFound))

		sum := 0
		for _, num := range stats.NumbersFound {
			sum += num
		}
		fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(stats.NumbersFound)))
	}
}

type KeyValue struct {
	Key   string
	Value int
}

type KeyValueRune struct {
	Key   rune
	Value int
}

func getTopItems(m map[string]int, n int) []KeyValue {
	items := make([]KeyValue, 0, len(m))
	for k, v := range m {
		items = append(items, KeyValue{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	if len(items) > n {
		return items[:n]
	}
	return items
}

func getTopItemsRune(m map[rune]int, n int) []KeyValueRune {
	items := make([]KeyValueRune, 0, len(m))
	for k, v := range m {
		items = append(items, KeyValueRune{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	if len(items) > n {
		return items[:n]
	}
	return items
}

func main() {
	var text string

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		text = string(content)
	} else {
		fmt.Println("Enter text (press Ctrl+D when finished):")
		scanner := bufio.NewScanner(os.Stdin)
		var builder strings.Builder
		for scanner.Scan() {
			builder.WriteString(scanner.Text() + "\n")
		}
		text = builder.String()
	}

	stats := Analyze(text)
	stats.DisplayStats()
}
