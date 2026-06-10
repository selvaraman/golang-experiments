package main

import "fmt"

func removeDuplicates(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	writeAt := 1
	for i:=1; i<len(nums); i++ {
		if nums[i] != nums[i-1] {
			if writeAt != i {
				nums[writeAt] = nums[i]
			}
			writeAt++
		}
	}
	return nums[:writeAt]
}

func main() {
	nums := []int {1, 1, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
