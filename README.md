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