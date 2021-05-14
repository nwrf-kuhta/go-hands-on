package main

import "fmt"

// スライスを操作する
func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)

	a = push(a, 1000)
	fmt.Println(a)

	a = pop(a)
	fmt.Println(a)

	a = unshift(a, 1000)
	fmt.Println(a)

	a = shift(a)
	fmt.Println(a)

	a = insert(a, 1000, 2)
	fmt.Println(a)

	a = remove(a, 2)
	fmt.Println(a)
}

// pushは、スライスの最後に追加
func push(a []int, v int) []int {
	return append(a, v)
}

// popは、スライスの最後を削除
func pop(a []int) []int {
	return a[:len(a)-1]
}

// unshiftは、スライスの最初に追加
func unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

// shiftは、スライスの最初を削除
func shift(a []int) []int {
	return a[1:]
}

// insertは、スライスの指定の位置に追加
func insert(a []int, v int, p int) []int {
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

// removeは、スライスの指定の位置を削除
func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
