package main



func merge(nums1 []int, m int, nums2 []int, n int)  {
    insertPoint := m + n - 1

    for m > 0 && n > 0 {
        if nums1[m - 1] > nums2[n - 1] {
            nums1[insertPoint] = nums1[m - 1]
           m--
        } else {
            nums1[insertPoint] = nums2[n - 1]
            n--
        }
        insertPoint--
    }

    // add missing elements from nums2
    for n > 0 {
        nums1[insertPoint] = nums2[n - 1]
        n--
        insertPoint--
    }

    for m > 0 {
        nums1[insertPoint] = nums1[m - 1] 
        m--
        insertPoint--
    }
}