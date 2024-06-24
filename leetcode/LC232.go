package leetcode

type MyQueue struct {
	stk1 []int
	stk2 []int
}

func QueueConstructor() MyQueue {
	return MyQueue{
		stk1: make([]int, 0),
		stk2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.stk1) > 0 {
		this.stk2 = append(this.stk2, this.Pop())
	}
	this.stk1 = append(this.stk1, x)
	for len(this.stk2) > 0 {
		n := len(this.stk2) - 1
		pop := this.stk2[n]
		this.stk2 = this.stk2[:n]
		this.stk1 = append(this.stk1, pop)
	}
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		panic("empty")
	}
	n := len(this.stk1) - 1
	pop := this.stk1[n]
	this.stk1 = this.stk1[:n]
	return pop
}

func (this *MyQueue) Peek() int {
	n := len(this.stk1) - 1
	return this.stk1[n]
}

func (this *MyQueue) Empty() bool {
	return len(this.stk1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func RunLC232() {
	obj := QueueConstructor()
	obj.Push(1)
	obj.Push(2)
	println(obj.Peek())
	println(obj.Pop())
	println(obj.Empty())
}
