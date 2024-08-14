package leetcode

type StockSpanner struct {
	Stack [][]int
}

func Constructor2() StockSpanner {
	return StockSpanner{
		Stack: make([][]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	val := 1
	for len(this.Stack) > 0 && this.Stack[len(this.Stack)-1][0] <= price {
		val += this.Stack[len(this.Stack)-1][1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, []int{price, val})
	return val
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
