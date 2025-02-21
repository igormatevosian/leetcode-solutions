package golang

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	median := (len(nums1) + len(nums2)) / 2
	index1, index2 := 0, 0
	res := 0
	if (len(nums1)+len(nums2))%2 == 0 {
		for i := 0; i <= median; i++ {
			if index1 != len(nums1) && (index2 == len(nums2) || nums1[index1] < nums2[index2]) {
				if i == median || i == median-1 {
					res += nums1[index1]
				}
				index1++
			} else {
				if i == median || i == median-1 {
					res += nums2[index2]
				}
				index2++
			}
		}
		return float64(res) / 2.0
	} else {
		for i := 0; i <= median; i++ {
			if index1 != len(nums1) && (index2 == len(nums2) || nums1[index1] < nums2[index2]) {
				if i == median {
					res = nums1[index1]
					break
				}
				index1++
			} else {
				if i == median {
					res = nums2[index2]
					break
				}
				index2++
			}
		}
		return float64(res)
	}
}
