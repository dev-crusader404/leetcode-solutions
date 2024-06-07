package leetcode

//   Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
	data []int
}

func (this *Iterator) hasNext() bool {
	return len(this.data) != 0
}

func (this *Iterator) next() int {
	n := len(this.data)
	pop := this.data[n-1]
	this.data = this.data[:n-1]
	return pop
}

type PeekingIterator struct {
	itr *Iterator
	top int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	pk := &PeekingIterator{
		itr: iter,
	}
	if iter.hasNext() {
		pk.top = iter.next()
	}
	return pk
}

func (this *PeekingIterator) hasNext() bool {
	return this.top != 0
}

func (this *PeekingIterator) next() int {
	pop := this.top
	if this.itr.hasNext() {
		this.top = this.itr.next()
	} else {
		this.top = 0
	}
	return pop
}

func (this *PeekingIterator) peek() int {
	return this.top
}
