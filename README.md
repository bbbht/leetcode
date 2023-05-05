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

### 动态规划

### 回溯

### 前缀和
