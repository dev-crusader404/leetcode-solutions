package leetcode

type OrderedStream struct {
	List []string
	idx  int
}

func Constructors(n int) OrderedStream {
	return OrderedStream{
		List: make([]string, n),
		idx:  0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.List[idKey-1] = value

	temp := []string{}
	for this.idx < len(this.List) && this.List[this.idx] != "" {
		temp = append(temp, this.List[this.idx])
		this.idx++
	}
	return temp
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
