package main

import "fmt"

// スライスの長さと容量
func main() {
	// スライスは配列の参照に過ぎない
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]
	fmt.Println(a) // 10, 20, 30, 40, 50
	fmt.Println(b) // 10, 20, 30
	a[0] = 100
	fmt.Println(a) // 100, 20, 30, 40, 50
	fmt.Println(b) // 100, 20, 30
	b[1] = 200
	fmt.Println(a) // 100, 200, 30, 40, 50
	fmt.Println(b) // 100, 200, 30
}
