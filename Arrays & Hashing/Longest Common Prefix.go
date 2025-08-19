package main

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {

	res := []byte{}

	for i := 0; i < len(strs[0]); i++ {

		for _, s := range strs {
			if i == len(s) || strs[0][i] != s[i] {
				return string(res)
			}
		}

		res = append(res, strs[0][i])
	}

	return string(res)

}
