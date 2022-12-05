package main

// 225 implement-stack-using-queues 2022-11-30 11:40:58

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.q1)-1; i++ {
		this.q2 = append(this.q2, this.q1[i])
	}
	res := this.q1[len(this.q1)-1]
	this.q1, this.q2 = this.q2, nil
	return res
}

func (this *MyStack) Top() int {
	res := this.Pop()
	this.q1 = append(this.q1, res)
	return res
}

func (this *MyStack) Empty() bool {
	if len(this.q1) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
