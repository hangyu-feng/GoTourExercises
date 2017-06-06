// 让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址
// 例如 IPAddr{1, 2, 3, 4} 应当输出 "1.2.3.4"
package main

import (
	"fmt"
	// "strings"
	"strings"
)

type IPAddr [4]byte

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

func (i IPAddr) String() string {
	return strings.Join(strings.Fields(strings.Trim(fmt.Sprint(i), "[]")), ".")
}
