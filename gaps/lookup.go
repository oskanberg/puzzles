package main

import "math/rand"

// LetterLookup is a lookup difference letter -> slice of words
type LetterLookup map[string][][]string

var letterLookup LetterLookup = make(map[string][][]string)

// LookupSample returns a random set of words that differ by 'delta'
func (l LetterLookup) LookupSample(delta string) []string {
	index := rand.Intn(len(l[delta]))
	return l[delta][index]
}

func (l LetterLookup) add(key string, words []string) {
	if val, ok := l[key]; ok {
		l[key] = append(val, words)
	} else {
		l[key] = make([][]string, 1)
		l[key][0] = words
	}
}

// AddToLookup adds a group of words to the lookup
func (l LetterLookup) AddToLookup(entry []string) {
	for i, char := range entry[0] {
		if byte(char) != entry[1][i] {
			// fmt.Println("difference between", entry[0], entry[1], "is", string(char), "/", string(entry[1][i]))
			l.add(string(char), entry)
			l.add(string(entry[1][i]), entry)
			break
		}
	}
}
