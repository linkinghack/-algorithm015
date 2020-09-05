package codes

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	groupMap := map[string][]string{}
	for _, str := range strs {
		charArray := strings.Split(str, "")
		sort.Strings(charArray)
		standardStr := strings.Join(charArray, "")

		groupMap[standardStr] = append(groupMap[standardStr], str)
	}

	for _, group := range groupMap {
		result = append(result, group)
	}
	return result
}

// @lc code=end
