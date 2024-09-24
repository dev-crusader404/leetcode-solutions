package leetcode

type ProductOfNumber struct {
	prefixProd []int
}

func Constr() ProductOfNumber {
	return ProductOfNumber{prefixProd: make([]int, 0)}
}

func (this *ProductOfNumber) Add(num int) {
	if len(this.prefixProd) == 0 {
		this.prefixProd = append(this.prefixProd, 1)
	}
	if num > 0 {
		lastProd := this.prefixProd[len(this.prefixProd)-1]
		this.prefixProd = append(this.prefixProd, lastProd*num)
	} else {
		this.prefixProd = this.prefixProd[:0]
		this.prefixProd = append(this.prefixProd, 1)
	}
}

func (this *ProductOfNumber) GetProduct(k int) int {
	n := len(this.prefixProd)
	if k < n {
		return this.prefixProd[n-1] / this.prefixProd[n-k-1]
	}
	return 0
}

/**
 * Your ProductOfNumber object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
