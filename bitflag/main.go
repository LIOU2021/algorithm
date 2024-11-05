package main

import (
	"fmt"
	"unsafe"
)

const (
	mode1 uint8 = 1 << iota // 0b00000001
	mode2                   // 0b00000010
	mode3                   // 0b00000100
	mode4                   // 0b00001000
	mode5                   // 0b00010000
	mode6                   // 0b00100000
	mode7                   // 0b01000000
	mode8                   // 0b10000000
)

func render(flag uint8) {
	if (flag & mode1) == mode1 {
		fmt.Println("mode 1")
	}

	if (flag & mode2) == mode2 {
		fmt.Println("mode 2")
	}

	if (flag & mode3) == mode3 {
		fmt.Println("mode 3")
	}

	if (flag & mode4) == mode4 {
		fmt.Println("mode 4")
	}

	if (flag & mode5) == mode5 {
		fmt.Println("mode 5")
	}

	if (flag & mode6) == mode6 {
		fmt.Println("mode 6")
	}

	if (flag & mode7) == mode7 {
		fmt.Println("mode 7")
	}

	if (flag & mode8) == mode8 {
		fmt.Println("mode 8")
	}
}

func main() {
	// 为了节省记忆体, const用uint8资料型态, 只占据1 byte, golang 容量最小单位为1 byte
	fmt.Println("啟用模式 1, 2")
	render(mode1 | mode2)
	fmt.Println("啟用模式 1, 3, 4")
	render(mode1 | mode3 | mode4)
	fmt.Println("啟用模式 5, 7, 8")
	render(mode5 | mode7 | mode8)

	fmt.Println("=====================")

	fmt.Printf("Has(mode1|mode2, mode1): %t\n", Has(mode1|mode2, mode1))                         // 模拟发生了mode1跟mode2事件，接收到讯息后判别是否有mode1事件
	fmt.Printf("Has(mode3|mode8, mode8): %t\n", Has(mode3|mode8, mode8))                         // 模拟发生了mode3跟mode8事件，接收到讯息后判别是否有mode8事件
	fmt.Printf("Has(mode3|mode8, mode3): %t\n", Has(mode3|mode8, mode3))                         // 模拟发生了mode3跟mode8事件，接收到讯息后判别是否有mode3事件
	fmt.Printf("Has(mode3|mode8, mode4): %t\n", Has(mode3|mode8, mode4))                         // 模拟发生了mode3跟mode8事件，接收到讯息后判别是否有mode4事件
	fmt.Printf("Has(mode3|mode8|mode7, mode3|mode8): %t\n", Has(mode3|mode8|mode7, mode3|mode8)) // 模拟发生了mode3跟model7跟mode8事件，接收到讯息后判别是否有同时发生mode3跟mode8事件

	fmt.Println("=====================")

	var a int
	var b float64
	var c bool
	var d string
	var e [10]int
	var f int8
	var g uint8

	fmt.Printf("int: %d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("bool: %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("string: %d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("array of 10 ints: %d bytes\n", unsafe.Sizeof(e))
	fmt.Printf("mode1: %d bytes\n", unsafe.Sizeof(mode1))
	fmt.Printf("f: %d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("g: %d bytes\n", unsafe.Sizeof(g))
	fmt.Printf("mode5: %d bytes. %08b, %v\n", unsafe.Sizeof(mode5), mode5, 0b10000)

	fmt.Println(0b10000000, mode8)
}

// source: 模拟类似file system 或是异步传递来的msg, 可以一次传递多种类型进来
// has: 判断是否有特定类型数据
func Has(source, has uint8) bool {
	return source&has != 0
}
