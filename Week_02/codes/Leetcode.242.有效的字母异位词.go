package codes

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	return solveBySort(s, t)
}

func solveBySort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	sort.Strings(sArr)
	sort.Strings(tArr)

	for i := 0; i < len(sArr); i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
}

func solveWithHashTable(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alpha := make(map[int32]int)
	for _, c := range s {
		if _, exists := alpha[c]; exists {
			alpha[c]++
		} else {
			alpha[c] = 1
		}
	}

	for _, c := range t {
		if _, exists := alpha[c]; exists {
			alpha[c]--
		} else {
			// character doesn't exist
			return false
		}
	}

	for _, v := range alpha {
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
