package main

import (
	"math/rand"
)

type WordNode struct {
	word       string
	neighbours map[string]*WordNode
}

func (n *WordNode) AddNeighbour(node *WordNode) {
	n.neighbours[node.word] = node
}

func (n *WordNode) GetRandomNeighbour() *WordNode {
	keys := make([]string, len(n.neighbours))

	i := 0
	for k := range n.neighbours {
		keys[i] = k
		i++
	}

	i = rand.Intn(len(keys))
	return n.neighbours[keys[i]]
}

func NewWordNode(word string) *WordNode {
	return &WordNode{
		word:       word,
		neighbours: make(map[string]*WordNode),
	}
}

type WordWeb struct {
	index map[string]*WordNode
}

func (w *WordWeb) GetWordNode(word string) *WordNode {
	node, ok := w.index[word]
	if !ok {
		node = NewWordNode(word)
		w.index[word] = node
	}
	return node
}

func (w *WordWeb) AddWordPair(firstWord, secondWord string) {
	firstNode := w.GetWordNode(firstWord)
	secondNode := w.GetWordNode(secondWord)
	firstNode.AddNeighbour(secondNode)
	secondNode.AddNeighbour(firstNode)
}

func (w *WordWeb) GetRandomWord() *WordNode {
	keys := make([]string, len(w.index))

	i := 0
	for k := range w.index {
		keys[i] = k
		i++
	}

	i = rand.Intn(len(keys))
	return w.index[keys[i]]
}

func NewWordWeb() *WordWeb {
	return &WordWeb{
		index: make(map[string]*WordNode),
	}
}
