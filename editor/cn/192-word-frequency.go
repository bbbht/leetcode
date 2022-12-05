package main
// 统计词频 2022-07-29 17:12:38

//There is no code of Go type for this problem
`
// sed 's/ /\n/g' words.txt | grep -v '^\s*$' | sort | uniq -c | sort -rn |awk '{print $2,$1}'