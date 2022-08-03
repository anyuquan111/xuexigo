package main

import (
	"fmt"
	"sort"
)

// 自定义排序需要重写类型和swap 方法和less方法
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i int, j int) bool {
	return len(s[i]) < len(s[j])
}
func main() {
	fruit := []string{"aannnnnna", "bbbbb", "xccccccc"}
	sort.Sort(byLength(fruit))
	fmt.Println(fruit)
}
