package main

// PrefixSums builds a prefix sum array with a sentinel value at index 0.
// Input:  arr = [2, 4, -3, 5]
// Output: ps  = [0, 2, 6, 3, 8]
// Explanation: ps[i+1] = ps[i] + arr[i]

type RangeSums struct {
	ps []int
}

func ConstructorRangeSums(arr []int) RangeSums {
	return RangeSums{
		ps: PrefixSums(arr),
	}
}

func PrefixSums(arr []int) []int {
	n := len(arr)
	ps := make([]int, n+1) //p[0] = 0

	for i := 0; i < len(arr); i++ {
		ps[i+1] = ps[i] + arr[i]
	}

	return ps
}

// RangeSum returns the sum of arr[l..r] inclusive, using the prefix sums.
// arr: [2,4,-3,5], ps: [0,2,6,3,8]
// Example: sum(1..3) = ps[3+1] - ps[1] = ps[4] - ps[1] = 8 - 2 = 6
func (r *RangeSums) RangeSum(left, right int) int {
	return r.ps[right+1] - r.ps[left]
}
