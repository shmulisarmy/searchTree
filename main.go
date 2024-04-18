package main

import (
        "bufio"
        "os"
)

type SearchTree struct {
        root map[byte]*SearchTree
        word string
}

func insert(tree *SearchTree, word string) {
        node := tree
        var index int = 0
        for index < len(word) {
                letter := word[index]
                if node.root[letter] == nil {
                        node.root[letter] = &SearchTree{root: make(map[byte]*SearchTree)}
                }
                node = node.root[letter]
                index++
        }
        node.word = word
}

func isWord(tree *SearchTree, word string) bool {
        node := tree
        var index int = 0
        for index < len(word) {
                letter := word[index]
                if node.root[letter] == nil {
                        return false
                }
                node = node.root[letter]
                index++
        }
        return node.word == word
}
func isPrefix(tree *SearchTree, word string) bool {
        node := tree
        var index int = 0
        for index < len(word) {