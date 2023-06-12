package main


func TwoSum(nums []int, target int) []int {
	database := make(map[int]int)
	for i, num := range nums {
		complemento := target - num

		if existe, ok := database[complemento]; ok {
			return []int{existe, i}
		}
		database[num] = i
	}
	return []int{}
}

