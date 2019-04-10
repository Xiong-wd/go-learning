//在不断扩充切片的时候，len的值适合添加的元素保持一致，每次添加一个元素len的值会增加1，见代码17-20
//cap的值会每次乘以2的频率增加，当len的值和cap的值相等的时候，如果继续添加len的值，cap的值会翻倍，并且按照 1，2，4，8，16，32，64，128 ... 的规律



package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil
	//
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	// 初始化切片的方式，第一种需要指定默认的值
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	// 第二种不需要指定默认的值，但是可以使用make指定切片的长度
	s2 := make([]int, 16)
	// 第三种不仅可以指定切片的长度，还可以指定相对应的cap的值
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)




	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

