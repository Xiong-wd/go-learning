// 在函数式编程中，函数是一等公民，可以作为参数，变量，返回值，进行传递
// golang实现了函数的闭包，高阶函数

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

// 实现read接口可以读文件的时候调用
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	fmt.Println("_______")
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = Fibonacci()
	printFileContents(f)
}

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}