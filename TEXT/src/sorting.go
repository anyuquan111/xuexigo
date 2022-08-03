package main

import (
	"fmt"
	"sort"
)

// 自定义排序需要自定义累重写swap和less函数
func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

}
