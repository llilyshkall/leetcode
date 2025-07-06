package t1865

type FindSumPairs struct {
	m1, m2 map[int]int
	nums2  []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	ret := FindSumPairs{
		m1:    make(map[int]int),
		m2:    make(map[int]int),
		nums2: nums2,
	}
	for _, n := range nums1 {
		ret.m1[n]++
	}
	for _, n := range nums2 {
		ret.m2[n]++
	}
	return ret
}

func (this *FindSumPairs) Add(index int, val int) {
	n := this.nums2[index]
	this.m2[n]--
	this.m2[n+val]++
	this.nums2[index] += val
}

func (this *FindSumPairs) Count(tot int) int {
	ret := 0
	for k, v := range this.m1 {
		ret += v * this.m2[tot-k]
	}
	return ret
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
