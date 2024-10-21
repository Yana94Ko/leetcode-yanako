func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i:=m; i<m+n; i++ {
        nums1[i] = nums2[i-m]
    }
    sort.Slice(nums1, func(j, k int) bool {
        return nums1[j] < nums1[k]
    })
}