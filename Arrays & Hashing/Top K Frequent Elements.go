package main

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	frequency := make([][]int, len(nums))

	for _, num := range nums {
		count[num]++
	}

	for num, freq := range count {
		frequency[freq-1] = append(frequency[freq-1], num)
	}

	var res []int
	for i := len(frequency) - 1; i >= 0; i-- {
		for _, num := range frequency[i] {
			res = append(res, num)

			if len(res) == k {
				return res
			}
		}
	}

	return res

}
