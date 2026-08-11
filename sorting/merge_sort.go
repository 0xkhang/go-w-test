package sorting

func MergeSorted(nums1, nums2 []int) []int {
	result := make([]int, 0)
	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	for j < len(nums1) {
		result = append(result, nums2[j])
		j++
	}

	return result
}
