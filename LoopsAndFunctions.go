//练习：循环和函数
//作为练习函数和循环的简单途径，用牛顿法实现开方函数。

//在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值：
//z = (z^2-x)/2z

//为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，使得当值停止改变（或改变
//非常小）的时候退出循环。观察迭代次数是否变化。结果与 math.Sqrt 接近吗？

package main

import "fmt"

func main() {
	fmt.Println(Sqrt(2))
}

func Sqrt(x float64) float64 {
	result := x
	for i:=0; i<10; i++ {
		result = result - ((result * result - x)  / (2 * result))
	}
	return result
}
