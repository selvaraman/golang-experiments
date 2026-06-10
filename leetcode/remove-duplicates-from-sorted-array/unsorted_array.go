package main

import "fmt"

func removeDuplicates(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	seen := make(map[int]bool)
	seen[nums[0]] = true
	writeAt := 1
	for i:=1; i<len(nums); i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			nums[writeAt] = nums[i]
			writeAt++
		}
	}
	return nums[:writeAt]
}
func main() {
	//nums := []int {1, 2, 1, 3}
	nums := []int {1, 2, 1, 3, 2, 4, 4 }
	fmt.Println(removeDuplicates(nums))
}
