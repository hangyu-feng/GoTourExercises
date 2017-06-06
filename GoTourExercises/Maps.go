//实现 WordCount。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。

//你会发现 strings.Fields 很有帮助。

package main

import (
	"strings"
)

func main() {}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	l := strings.Fields(s)
	for _, i := range l {
		m[i] += 1
	}
	return m
}
