package main

import "fmt"

func main() {

	//neetcode -1
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	//neetcode-2
	fmt.Println(getConcatenation([]int{1, 2, 1}))
	fmt.Println(getConcatenation([]int{1, 3, 2, 1}))

	//netcode-3
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))

	//neetcode-4
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

	//neetcode-5
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))

	//neetcode-6
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))

	//neetcode-7
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

	//neetcode-8
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

	//neetcode-9
	myHashSet := Constructor()
	myHashSet.Add(1)      // set = [1]
	myHashSet.Add(2)      // set = [1, 2]
	myHashSet.Contains(1) // return True
	myHashSet.Contains(3) // return False, (not found)
	myHashSet.Add(2)      // set = [1, 2]
	myHashSet.Contains(2) // return True
	myHashSet.Remove(2)   // set = [1]
	myHashSet.Contains(2) // return False, (already removed)

}
