package main

import (
	"bufio"
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

	var wordWeb = NewWordWeb()

	diffWordsFeed := createFeed("../resources/1dwords.csv")
	for words := range diffWordsFeed {
		for _, source := range words {
			for _, target := range words {
				if source == target {
					continue
				}
				wordWeb.AddWordPair(source, target)
			}
		}
	}

	steps := flag.Int("steps", 4, "steps in the ladder")
	flag.Parse()

	word := wordWeb.GetRandomWord()
	fmt.Println(word.word)

	for step := 0; step < *steps; step++ {
		word = word.GetRandomNeighbour()
		fmt.Println(word.word)
	}

}
