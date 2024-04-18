package main

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
		letter := word[index]
		if node.root[letter] == nil {
			return false
		}
		node = node.root[letter]
		index++
	}
	return true
}

func firstNResults(tree *SearchTree, letters string) []string {
	node := tree
	var index int = 0
	for index < len(letters) {
		letter := letters[index]
		if node.root[letter] != nil {
			node = node.root[letter]
		}
		index++
	}

	var found []string = []string{}

	var que []*SearchTree = []*SearchTree{node}

	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node.word != "" {
			found = append(found, node.word)
			if len(found) >= 3 {
				return found
			}
		}
		for _, child := range node.root {
			que = append(que, child)
		}
	}
	return found
}

func insertList(tree *SearchTree, words []string) {
	var i int = 0

	for i < len(words) {
		insert(tree, words[i])
		i++
	}
}
