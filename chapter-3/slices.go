// 数组的值是固定的，并且，不同长度的数组属于不同的类型
// 数组是值传递
// 切片（slices）是数组的一个映射，底层是数组，切片包含三个部分，ptr表示起始部分，len表示切片的长度，cap表示切片的容量
// 切片的容量表示切片取的底层数组向后未展示的部分，切片可以向后拓展，但是不能向前拓展
// 切片的传递和修改是引用传递，并且切片的切片的修改同样是指向最底层数组，并且会改变以此数组为基础的额所有的切片的值
// 当前切片不可超过自身len的长度，向后拓展的切片不可超过底层数组的cap
// 切片可以理解为数组的一个视图
// 当添加元素大于cap的值，系统会重新分配更大的底层数组
// 当使用append向一个数组里面添加元素的时候，都是值传递，所以必须再以一个值来接受新的数组

package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)

	// Uncomment to run sliceOps demo.
	// If we see undefined: sliceOps
	// please try go run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
	// sliceOps()
}
