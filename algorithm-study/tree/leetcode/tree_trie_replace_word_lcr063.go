package leetcode

import "strings"

// https://leetcode.cn/problems/UhWRSj/

func replaceWords(dictionary []string, sentence string) string {
	if len(dictionary) == 0 || sentence == "" {
		return sentence
	}

	dictTrie := &Trie{}
	for _, dict := range dictionary {
		dictTrie.Insert(dict)
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = Replace(*dictTrie, word)
	}

	return strings.Join(words, " ")
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (root *Trie) Insert(s string) {
	cur := root
	for _, c := range s {
		pathIndex := c - 'a'
		if cur.next[pathIndex] == nil {
			cur.next[pathIndex] = &Trie{}
		}
		cur = cur.next[pathIndex]
	}
	cur.isEnd = true
}

func Replace(dictTrie Trie, word string) string {
	cur := dictTrie
	for i, c := range word {
		pathIndex := c - 'a'
		if cur.next[pathIndex] == nil {
			return word
		}

		cur = *cur.next[pathIndex]

		if cur.isEnd {
			return word[:i+1]
		}
	}
	return word
}
