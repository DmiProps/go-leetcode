package problems

import "fmt"

// Problem127 is runs solution of problem and shows result.
func Problem127() {

	// variant 1: ans = 5
	beginWord := "hit"
	endWord := "cog"
	wordList := []string {"hot", "dot", "dog", "lot", "log", "cog"}

	// variant 2: ans = 2
	/*beginWord := "a"
	endWord := "c"
	wordList := []string {"a", "b", "c"}*/

	fmt.Println(ladderLength(beginWord, endWord, wordList))

}

// Problem #127 (medium level): Word Ladder.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	
	endIndex := -1;

	transMap := make([][]int, len(wordList), len(wordList))
	work	 := make([][2]int, 0, len(wordList))
	used	 := make([]bool, len(wordList), len(wordList))

	for i, s := range wordList {
		if transMap[i] == nil {
			transMap[i] = make([]int, 0)
		}
		if s == beginWord {
			continue
		}
		if s == endWord {
			endIndex = i
		}
		if isTransform(beginWord, s) {
			if endIndex == i {
				return 2
			}
			work = append(work, [2]int{i, 2})
			used[i] = true
		}
		for j := i + 1; j < len(wordList); j++ {
			if isTransform(s, wordList[j]) {
				transMap[i] = append(transMap[i], j)
				if transMap[j] == nil {
					transMap[j] = make([]int, 0)
				}
				transMap[j] = append(transMap[j], i)
			}
		}
	}

	if endIndex < 0 {
		return 0
	}
	
	list := []int{};
	cur := 0
	for len(work) > cur {
		list = transMap[work[cur][0]]
		for _, i := range list {
			if used[i] {
				continue
			}
			if i == endIndex {
				return work[cur][1] + 1;
			}
			work = append(work, [2]int{i, work[cur][1] + 1})
			used[i] = true
		}
		cur++
	}
	return 0

}

func isTransform(from string, to string) bool {

	f := false;
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			if f {
				return false;
			}
			f = true
		}
	}
	return true

}