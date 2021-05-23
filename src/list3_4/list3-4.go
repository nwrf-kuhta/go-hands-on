package main

import "fmt"

// ポインタを引数に使う
func main() {
	n := 123
	fmt.Printf("value: %d. \n", n)
	change1(n)
	fmt.Printf("value: %d. \n", n) // 123
	change2(&n)
	fmt.Printf("value: %d. \n", n) // 246
}

// 値渡し
func change1(n int) {
	n *= 2
}

// 参照渡し（ポインタ渡し）
func change2(n *int) {
	*n *= 2
}
