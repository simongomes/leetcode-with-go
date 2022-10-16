func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		rest := target - nums[i]

		if _, key := hash[rest]; key {
			result := []int { hash[rest], i }
			return result
		}

		hash[nums[i]] = i
	}

	return []int{}

}
