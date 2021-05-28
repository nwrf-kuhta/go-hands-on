package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

// チャンネルで値を渡す
func main() {
	c := make(chan int)
	go total(100, c)
	fmt.Println("total: ", <-c)
	// チャンネルから値を取得する場合、その値が送られてくるまで処理を待つ
}
