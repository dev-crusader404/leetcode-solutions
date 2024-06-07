package leetcode

/*   Below is the interface for Iterator, which is already defined for you.

   type Iterator struct {

   }

   func (this *Iterator) hasNext() bool {
		// Returns true if the iteration has more elements.
   }

   func (this *Iterator) next() int {
		// Returns the next element in the iteration.
   }
*/

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
