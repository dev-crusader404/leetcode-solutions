package leetcode

type ProductOfNumbers struct {
	nums []int
}

func NewProductOfNumbers() ProductOfNumbers {
	return ProductOfNumbers{nums: make([]int, 0)}
}

func (this *ProductOfNumbers) Add(num int) {
	this.nums = append(this.nums, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	prod, i := 1, 0

	for k > 0 {
		prod *= this.nums[len(this.nums)-1-i]
		k--
		i++
	}
	return prod
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := NewProductOfNumbers();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
