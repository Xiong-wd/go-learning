// defer 会确保函数在程序结束的时候发生，当程序遇到defer的时候会先不执行，直到程序结束的时候再执行
// 如果有多个defer，defer会以一种压栈的方法，先进后出的原则执行
//

package main

import "fmt"

func tryDefer () {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

func main() {
	tryDefer()
}