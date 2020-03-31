package trie

import (
    "errors"
)

type trieNode struct {
    key byte
    finalNode bool
    children map[byte]*trieNode
}

type Trie struct {
    root *trieNode
}

func New() *Trie {
    return &Trie{root: &trieNode{finalNode: false, children: map[byte]*trieNode{}}}
}

func (trie *Trie) Insert(word string) error {
    if trie == nil || trie.root == nil {
        return errors.New("cannot insert a word in a nil trie")
    }
    if word == "" {
        return errors.New("cannot insert an empty word")
    }
    wordBytes := []byte(word)
    currentNode := trie.root
    for _, char := range wordBytes {
        nextNode, ok := currentNode.children[char]
        if !ok {
            nextNode = &trieNode{key: char, finalNode: false, children: map[byte]*trieNode{}}
            currentNode.children[char] = nextNode
        }
        currentNode = nextNode
    }
    currentNode.finalNode = true
    return nil
}

func (trie *Trie) Search(word string) (bool, error) {
    if trie == nil || trie.root == nil {
        return false, errors.New("cannot search a word in a nil trie")
    }
    if word == "" {
        return false, nil
    }
    wordBytes := []byte(word)
    currentNode := trie.root
    for _, char := range wordBytes {
        nextNode, ok := currentNode.children[char]
        if !ok {
            return false, nil
        }
        currentNode = nextNode
    }
    return currentNode.finalNode, nil
}
