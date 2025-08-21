package main

import "strconv"

// Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

// Please implement encode and decode

// Example 1:

// Input: ["neet","code","love","you"]

// Output:["neet","code","love","you"]
// Example 2:

// Input: ["we","say",":","yes"]

// Output: ["we","say",":","yes"]
// Constraints:

// 0 <= strs.length < 100
// 0 <= strs[i].length < 200
// strs[i] contains only UTF-8 characters.

type Solution struct{}

func (s *Solution) encode(strings []string) string {
	if len(strings) <= 0 {
		return ""
	}
	var res string
	for _, s := range strings {
		res += strconv.Itoa(len(s)) + "#" + s
	}
	return res
}

func (s *Solution) decode(str string) []string {

	if str == "" {
		return []string{}
	}

	i := 0
	var res []string
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(str[i:j])

		i = j + 1
		res = append(res, str[i:i+length])
		i = i + length
	}

	return res

}
