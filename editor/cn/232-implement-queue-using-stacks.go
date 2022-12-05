package main

// 232 implement-queue-using-stacks 2022-11-30 16:19:10

// leetcode submit region begin(Prohibit modification and deletion)
// stack defines Stack1 stack
type stack232 struct {
	val []int
}

func myStack232() stack232  {
	return stack232{[]int{}}
}

func (this *stack232) Push(x int) {
	tmp := make([]int, 1, len(this.val)+1)
	tmp[0] = x
	this.val = append(tmp, this.val...)
}

func (this *stack232) Pop() int {
	tmp := this.val[0]
	this.val = this.val[1:]
	return tmp
}

func (this *stack232) Peek() int {
	return this.val[0]
}
func (this *stack232) Size() int {
	return len(this.val)
}

func (this *stack232) Empty() bool {
	if 0 == len(this.val) {
		return true
	}

	return false
}

type MyQueue struct {
	stack1 stack232
	stack2 stack232
}

func Constructor() MyQueue {
	var queue MyQueue
	queue.stack1 = myStack232()
	queue.stack2 = myStack232()

	return queue
}


func (this *MyQueue) Push(x int)  {
	this.stack1.Push(x)
}


func (this *MyQueue) Pop() int {
	if !this.stack2.Empty() {
		return this.stack2.Pop()
	} else {
		for ;!this.stack1.Empty(); {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
		return this.stack2.Pop()
	}
}

func (this *MyQueue) Peek() int {
	if !this.stack2.Empty() {
		return this.stack2.Peek()
	} else {
		for ;!this.stack1.Empty(); {
			tmp := this.stack1.Pop()
			this.stack2.Push(tmp)
		}
		return this.stack2.Peek()
	}
}

func (this *MyQueue) Empty() bool {
	return this.stack1.Empty() && this.stack2.Empty()
}
/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
