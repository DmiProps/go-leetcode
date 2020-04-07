package problems

import (
	"fmt"
	"strings"
)

func ProblemY() {

	obj := Constructor()
	obj.Insert("apple")
	param_2 := obj.Search("apple")
	param_3 := obj.StartsWith("app")
	fmt.Println(param_2)
	fmt.Println(param_3)

}

type Trie struct {
	M map[string]bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := new(Trie)
	trie.M = make(map[string]bool)
	return *trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.M[word] = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ok, _ := this.M[word]
	return ok
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for word, _ := range this.M {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
