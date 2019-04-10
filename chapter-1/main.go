package main

import (
	"errors"
	"fmt"
)

func f1() {
	var err error

	defer fmt.Println(err)

	err = errors.New("defer error")
	return
}

func f2() {
	var err error

	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer error")
	return
}

func f3() {
	var err error

	defer func(err error) {
		fmt.Println(err)
	}(err)

	err = errors.New("defer error")
	return
}

func main() {
	f1()
	f2()
	f3()
	fmt.Println(6 >> 1)

	fmt.Println(makechan(0,0))
}

// 初始化 channel
func makechan(t int, size int) int{
	var a int
	// 可以学习一下这种空 switch 的写法，比 if else 好看一些
	switch {
	case size > 0 :
		fmt.Println("size")
		a = size
	case t > 0:
		fmt.Println("t")
		a = t
	default:
		fmt.Println("default")
		a= 0
	}
	return a
}
