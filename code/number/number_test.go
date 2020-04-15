package number

import (
	"../selfutil"
	"math"
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

/*4. 寻找两个有序数组的中位数
* 找出中位数实际则是找出第k个数，分别在两个数组里面找出第k/2个数A和B，删除A和B较小的那个数组的前k/2的数。
理由是较小的那个数的位置最大也就是k-1,因为后面最少还有一个比他大的数，
*/
func TestFindMedian(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{2,3,4,5,6}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	if sum % 2 == 0 {
		return (findK(nums1, 0, nums2,0, sum) + findK(nums1, 0, nums2, 0, sum / 2 + 1)) / 2
	}
	return findK(nums1, 0, nums2, 0, sum/2 + 1)
}
func findK(nums1 []int, left1 int, nums2 []int, left2 int, k int) float64 {
	if left1 >= len(nums1) {
		return float64(nums2[left2+k-1])
	}
	if left2 >= len(nums2) {
		return float64(nums1[left1+k-1])
	}
	if k == 1 {
		return float64(selfutil.Min(nums1[left1], nums2[left2]))
	}
	mid1, mid2 := nums1[left1 + k / 2 -1], nums2[left2 + k / 2 -1]
	if left1+k/2-1 > len(nums1) {
		mid1 = math.MaxInt64
	}

	if left1+k/2-1 > len(nums1) {
		mid1 = math.MaxInt64
	}
	if mid1 <= mid2 {
		return findK(nums1, left1 + k /2, nums2, left2, k - k / 2)
	}
	return findK(nums1, left1, nums2, left2 + k / 2, k - k / 2)
}
