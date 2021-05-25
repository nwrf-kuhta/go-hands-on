package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

// チャンネルへの複数値の追加
func main() {
	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	x, y, z := <-c, <-c, <-c

	fmt.Println(x, y, z)
	// チャンネルは、単純に一つの値を保管するわけではなく、複数の値を保管することができる。
	// これは、FIFO（先入先出し）となっており、追加したものから順に値が取り出されるようになっている。
	// チャンネルから取り出される順序は、Goルーチンを呼び出した順とは限らないことに注意
}
