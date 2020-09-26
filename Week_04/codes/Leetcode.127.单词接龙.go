package codes

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	type BfsElem struct {
		word  string
		level int // 用于记录BFS层数，也是转换序列长度
	}

	wordPatternMap := map[string][]string{}
	L := len(beginWord)

	// 构建字典匹配表
	for _, word := range wordList {
		for i := 0; i < L; i++ {
			// hit -> h*t, *ht, hi*
			pattern := word[:i] + "*" + word[i+1:]
			wordPatternMap[pattern] = append(wordPatternMap[pattern], word)
		}
	}

	// BFS
	// 利用广度优先搜索同一层状态消费的步数相等
	queue := list.New()
	queue.PushBack(&BfsElem{beginWord, 1})
	visited := map[string]bool{}
	for queue.Len() > 0 {
		currentWordElem := queue.Front()
		queue.Remove(currentWordElem)
		currentWord := currentWordElem.Value.(*BfsElem)

		// 搜索过程即对每个word的每种pattern搜索，查看是否存在目标word
		for i := 0; i < len(currentWord.word); i++ {
			pattern := currentWord.word[:i] + "*" + currentWord.word[i+1:]
			for _, word := range wordPatternMap[pattern] {
				if word == endWord {
					return currentWord.level + 1 // 转换序列长度
				}

				if !visited[word] {
					visited[word] = true
					queue.PushBack(&BfsElem{word: word, level: currentWord.level + 1})
				}
			}
		}
	}
	return 0
}
