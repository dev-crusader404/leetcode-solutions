package leetcode

type LUPrefix struct {
	List []int
	idx  int
}

func Constructor3(n int) LUPrefix {
	return LUPrefix{
		List: make([]int, n+1),
		idx:  0,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.List[video-1] = video
	for len(this.List) > 0 && this.List[this.idx] != 0 {
		this.idx++
	}
}

func (this *LUPrefix) Longest() int {
	return this.idx
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
