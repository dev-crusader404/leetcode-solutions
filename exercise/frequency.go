package exercise

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := FreqMap{}
	for _, txt := range texts {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			tempMap := Frequency(s)
			mu.Lock()
			for k, v := range tempMap {
				result[k] += v
			}
			mu.Unlock()
		}(txt)
	}
	wg.Wait()
	return result
}
