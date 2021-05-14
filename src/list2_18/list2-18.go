package main

import "fmt"

// 要素を追加する
func main() {
	// スライスにappendで値を追加すると、元の配列の要素が書き変わっていく
	// 元の要素数以上にappendで追加すると、配列は変わらずスライスにだけ値が追加されていく
	a := [3]int{10, 20, 30}
	b := a[0:2]
	fmt.Println(a) // 10, 20, 30
	fmt.Println(b) // 10, 20
	b = append(b, 1000)
	fmt.Println(a) // 10, 20, 1000
	fmt.Println(b) // 10, 20, 1000
	b = append(b, 1000)
	fmt.Println(a) // 10, 20, 1000
	fmt.Println(b) // 10, 20, 1000, 1000
}
