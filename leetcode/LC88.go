package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	mLen, nLen, total := m-1, n-1, m+n-1

	for mLen >= 0 && nLen >= 0 {
		if nums1[mLen] > nums2[nLen] {
			nums1[total] = nums1[mLen]
			mLen--
		} else {
			nums1[total] = nums2[nLen]
			nLen--
		}
		total--
	}

	for nLen >= 0 {
		nums1[total] = nums2[nLen]
		nLen--
		total--
	}
}
