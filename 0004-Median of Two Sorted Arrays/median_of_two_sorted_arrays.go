func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findKth(nums1 []int, nums2 []int, start1 int, start2 int, k int) int {
	//if only one array is available for search
	if start1 >= len(nums1) {
		return nums2[start2+k-1]
	} else if start2 >= len(nums2) {
		return nums1[start1+k-1]
	}

	//if it is searching the first item
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	med1, med2, key1, key2 := start1+k/2, start2+k/2, 65535, 65535
	if med1 <= len(nums1) {
		key1 = nums1[med1-1]
	}

	if med2 <= len(nums2) {
		key2 = nums2[med2-1]
	}

	//binary search
	if key1 < key2 {
		return findKth(nums1, nums2, med1, start2, k-k/2)
	} else {
		return findKth(nums1, nums2, start1, med2, k-k/2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := len(nums1) + len(nums2)
	if num%2 == 1 {
		med := findKth(nums1, nums2, 0, 0, num/2+1)
		return float64(med) / 1.0
	} else {
		med1 := findKth(nums1, nums2, 0, 0, num/2+1)
		med2 := findKth(nums1, nums2, 0, 0, num/2)
		return float64(med1+med2) / 2.0
	}
}