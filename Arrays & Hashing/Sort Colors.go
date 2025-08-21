package main

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

// You must solve this problem without using the library's sort function.

// Example 1:

// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Example 2:

// Input: nums = [2,0,1]
// Output: [0,1,2]

// Constraints:

// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.

// Follow up: Could you come up with a one-pass algorithm using only constant extra space?
//one pass means we need to use one loop in whole function,
// other thing to note here is we need do perform sorting in in-place which means we are not allowed to create another int slice

func sortColors(nums []int) {

	low, mid := 0, 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}

}
