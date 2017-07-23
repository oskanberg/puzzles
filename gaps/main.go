package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func createFeed(path string) chan []string {
	ch := make(chan []string)
	go func() {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var line string
		for scanner.Scan() {
			line = scanner.Text()
			words := strings.Split(line, ",")
			ch <- words
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		close(ch)
	}()
	return ch
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	diffWordsFeed := createFeed("../resources/1dwords.csv")
	for words := range diffWordsFeed {
		letterLookup.AddToLookup(words)
	}

	englishWordsFeed := createFeed("../resources/englishwords.txt")
	var englishWords []string
	for words := range englishWordsFeed {
		englishWords = append(englishWords, words[0])
	}

	fmt.Println(englishWords)

	b, _ := json.MarshalIndent(letterLookup, "", "  ")
	fmt.Print(string(b))

	target := flag.String("target", "puzzle", "the final puzzle word")
	flag.Parse()

	for _, char := range *target {
		fmt.Println(string(char))
		sample := letterLookup.LookupSample(string(char))
		fmt.Println(sample)
	}
}
