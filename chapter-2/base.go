// 在go语言中，没有long类型，长整型用int64/int32表示，并且前面可以加u，例如 uint
// uint32 / uint64 表示无符号的整形，其中uintptr 表示指针类型
// golang中的char用rune表示，rune占四个字节32位，为了适应全球化，表达多国语言。
// 浮点型包括 float32 /float64 其中后者表示长浮点性
// complex64 / complex128 表示 复数 详见euler（）

// 复数 ： 我们把形如z=a+bi（a,b均为实数）的数称为复数，其中a称为实部，
// b称为虚部，i称为虚数单位。当z的虚部等于零时，常称z为实数；当z的虚部不等于零时，
// 实部等于零时，常称z为纯虚数。复数域是实数域的代数闭包，即任何复系数多项式在复数域中总有根。
// 复数是由意大利米兰学者卡当在十六世纪首次引入，经过达朗贝尔、棣莫弗、欧拉、高斯等人的工作，
// 此概念逐渐为数学家所接受。

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	euler()
	enum()
}


func euler() {
	s := 3 + 4i

	fmt.Println(cmplx.Abs(s))
	// 结果无限接近于0，浮点数的精度不准确
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi))

}


func enum() {
	const (
		b = 2 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b,kb,mb,gb,tb,pb)


}