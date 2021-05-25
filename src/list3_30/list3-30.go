package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

// mainからチャンネルに値が送れない？
func main() {
	c := make(chan int)
	c <- 100
	go total(c)
	time.Sleep(100 * time.Millisecond)
	// これはGoルーチンがデッドロックするため、fatal errorとエラーが発生しプログラムが強制終了する。
	// チャンネルは複数のスレッド間で値をやり取りするためのものなので、値を送る側と受け取る側の双方向で値の準備が整えってなければならない。
	// つまり、Goルーチンによるスレッドを実行した後でないとチャンネルは使えない。
}
