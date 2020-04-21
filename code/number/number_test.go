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
二分法求解，时间复杂度为log（m+n）
1、找出中位数实际则是找出第k个数，分别在两个数组里面找出第k/2个数A和B
2、删除A和B较小的那个数组的前k/2的数。理由是较小的那个数的位置最大也就是k-1,因为后面最少还有一个比他大的数
3、使用case来聚合多个条件运算，耗时从36ms降到12ms，代码也简洁不少
*/
func TestFindMedian(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	if sum % 2 == 0 {
		return (findK(nums1, 0, nums2,0, sum / 2) + findK(nums1, 0, nums2, 0, sum / 2 + 1)) / 2
	}
	return findK(nums1, 0, nums2, 0, sum/2 + 1)
}
func findK(nums1 []int, left1 int, nums2 []int, left2 int, k int) float64 {
	mid1, mid2 := 0, 0
	switch  {
	case left1 >= len(nums1) :
		return float64(nums2[left2+k-1])
	case left2 >= len(nums2):
		return float64(nums1[left1+k-1])
	case k == 1:
		return float64(selfutil.Min(nums1[left1], nums2[left2]))
	case left1 + k / 2 -1 >= len(nums1):
		mid1 = math.MaxInt64
	case left2 + k / 2 - 1 >= len(nums2):
		mid2 = math.MaxInt64
	default:
		mid1 = nums1[left1 + k / 2 -1]
		mid2 = nums2[left2 + k / 2 -1]
	}
	if mid1 <= mid2 {
		return findK(nums1, left1 + k / 2, nums2, left2, k - k / 2)
	}
	return findK(nums1, left1, nums2, left2 + k / 2, k - k / 2)
}

func TestReverse(t *testing.T) {
	println(math.MinInt32)
	println(math.MaxInt32)
	println(reverse(15342364690))

}
//7. 整数反转
func reverse(x int) int {
	y := 0
	if x == 0 {
		return  x
	}
	for x <= -1 || x >= 1  {
		y = y * 10 + x % 10
		x /= 10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}
