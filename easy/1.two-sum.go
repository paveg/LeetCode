package code

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// // Time complexity: O(n^2)
	// // Space complexity: O(1)
	// result := make([]int, 0, 2)
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			result = append(result, i, j)
	// 			return result
	// 		}
	// 	}
	// }
	// return result

	// Optimized version using a hash map
	// Time complexity: O(n)
	// Space complexity: O(n)
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		want := target - v
		fmt.Println("want:", want)
		if j, ok := seen[want]; ok {
			return []int{j, i}
		}
		seen[v] = i
		fmt.Println("seen:", seen)
	}
	return nil
}

// @lc code=end
