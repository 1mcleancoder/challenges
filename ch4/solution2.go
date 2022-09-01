package ch4

type solution2 struct {

}

func (s solution2) findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge arrays
	m := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))

	var i, j, k int
	for ; i < len(nums1) && j < len(nums2); k++ {

		if nums1[i] <= nums2[j] {
			m[k] = nums1[i]
			i++
		} else {
			m[k] = nums2[j]
			j++
		}
	}

	// add rem nums1
	for ; i < len(nums1); i++ {
		m[k] = nums1[i]
		k++
	}

	// add rem nums2
	for ; j < len(nums2); j++ {
		m[k] = nums2[j]
		k++
	}

	// calc median
	var med float64

	n := len(m) / 2
	if len(m) % 2 == 0 {
		med = float64(m[n-1] + m[n]) / float64(2)
	} else {
		med = float64(m[n])
	}

	return med
}
