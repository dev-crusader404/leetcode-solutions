package leetcode

type RecentCounter struct {
	counter []int
	start   int
	end     int
}

func Constructor933() RecentCounter {
	return RecentCounter{
		counter: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	for this.start < this.end && t-this.counter[this.start] > 3000 {
		this.start++
	}
	this.counter = append(this.counter, t)
	this.end++
	return this.end - this.start
}
