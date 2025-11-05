package code

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := [][]int{{}}

	// iterate through each number in nums
	result[0] = []int{} // [[]]
	for _, num := range nums {
		n := len(result)
		for i := 0; i < n; i++ {
			// create a new combination by adding the current num to each existing subset
			combination := make([]int, len(result[i]))
			// copy the existing subset
			copy(combination, result[i])
			// add the current num
			combination = append(combination, num)
			result = append(result, combination)
		}
	}
	return result
}

// @lc code=end
