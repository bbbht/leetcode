package main

import (
	"fmt"
)

// 303 range-sum-query-immutable 2022-12-02 15:05:02

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	sumArr []int
}


func Constructor(nums []int) NumArray {
	sumArr := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		sumArr[i+1] = nums[i]+sumArr[i]
	}
	fmt.Println(sumArr)
	return NumArray{sumArr: sumArr}
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.sumArr[right+1] - this.sumArr[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}