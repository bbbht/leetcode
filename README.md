# leetcode
偶尔做一做，要保持手感

## 环境
`Goland`中配置`LeetCode Editor`插件，自定义模板  
Code FileName:
```
$!{question.frontendQuestionId}-${question.titleSlug}
```
Code Template:
```
package main
// ${question.frontendQuestionId} ${question.titleSlug} $!velocityTool.date()

${question.code}

func main() {

}
```

## 总结
### 链表
#### 链表中点
依据快慢指针原理，快指针结束时，慢指针到中点
- 循环条件为 ``fast.next and fast.next``
	- 节点个数为奇数时，slow节点落在中间节点上
	- 节点个数为偶数数时，slow节点落在中间节点的左边一侧上
- 循环条件为 ``fast.next and fast.next.next``
	- 节点个数为奇数时，slow节点落在中间节点上
	- 节点个数为偶数数时，slow节点落在中间节点的右边一侧上
### 二分查找
[链接](https://labuladong.github.io/algo/di-ling-zh-bfe1b/wo-xie-le--3c789/)
#### 右边界的选择
区间开闭带来的区别  
1. 计算 mid 时要防止溢出，可以使用 mid = left + (right - left) >> 1 的写法
2. 如果 `right` 初始化为 `len(nums)`(数组长度)，
	1. 循环的判断条件为 `left < right`
	2. 循环中 `right` 的赋值表达式为 `right = mid` `left=mid+1`
3. 如果 `right` 初始化为 `len(nums) - 1`(索引最大值)
	1. 循环的判断条件为 `left <= right`
	2. 循环中 `right` 的赋值表达式为 `right = mid - 1` `left=mid+1`

## 位操作
### 异或操作
比如 136 只出现一次的数字，就利用以下三个规律，全部异或后得出
- 任何数和 0 做异或运算，结果仍然是原来的数
- 任何数和其自身做异或运算，结果是 0
- 异或运算满足交换律和结合律

### 二进制中1的个数
``n & (n-1)`` 的操作，可以将 n 对应的二进制最右侧的 1 置为 0

### 动态规划

### 回溯

### 前缀和
