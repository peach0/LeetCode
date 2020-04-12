package number

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	num := []int{2, 7, 11, 15}
	target := 9
	arr := twoSum(num, target)
	t.Error(arr)
}


func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) <= 0 {
		return result
	}
	dict := make(map[int]int)
	for k, v := range nums {
		if val, exist := dict[target-v]; exist {
			result = append(result, val)
			result = append(result, k)
		}
		dict[v] = k
	}
	return result
}

