package main

import "fmt"

func total(cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	cr <- t
}

// 双方向でやり取りするには
func main() {
	// 双方向でやり取りする場合、どちらのチャンネルがどちら向きに値を送るのかを把握しておく必要がある
	cs := make(chan int)
	cr := make(chan int)
	go total(cs, cr)
	cs <- 100
	fmt.Println("total:", <-cr)
}
