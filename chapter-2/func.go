package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)


//
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}
// 方法可以返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 方法可以作为参数传递，
func apply(op func(int, int) int, a, b int) int {
	// 通过反射获取方法的指针
	p := reflect.ValueOf(op).Pointer()
	// 运行时获取方法的名称
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 方法的参数，可以是多个值，用省略号代替
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
//
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	// 具体实现po方法，求出第一个数字的第二个数的次方
	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))
	//
	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
