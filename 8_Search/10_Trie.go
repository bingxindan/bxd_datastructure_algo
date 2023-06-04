package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func main() {
	word := "apple"
	prefix := "app"
	obj := Constructor()
	obj.Insert(word)
	response1 := obj.Search(word)
	fmt.Println(response1)
	response2 := obj.StartsWith(prefix)
	fmt.Println(response2)
	response3 := obj.SearchPrefix("a")
	fmt.Println(response3)
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}
