uses a search to quickly search for words that have been inserted into the tree

code sample:
```go

var tree SearchTree = SearchTree{root: make(map[byte]*SearchTree)}
insert(&tree, "hello")

var words []string = {"hello", "hey", "hi", "how", "you"}
insertList(tree, words)



var results []string = firstNResults(&tree, "k")
	for i := 0; i < len(results); i++ {
		println(results[i])
	}


```
