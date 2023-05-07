package main

// 155 min-stack 2023-05-07 20:27:54

// leetcode submit region begin(Prohibit modification and deletion)
// 两个栈，一个存储元素，一个存储对应的最小值

type MinStack struct {
	elements, min []int
	l             int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	if this.l == 0 {
		this.min = append(this.min, val)
	} else {
		min := this.GetMin()
		if val < min {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
