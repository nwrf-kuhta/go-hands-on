package main

import (
	"fmt"
	"strconv"
	"time"
)

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "second:["+<-c+"]")
	}
}

// スレッド間のチャンネルのやり取り
func main() {
	// 複数のスレッド間でチャンネルを利用する場合は、同じチャンネルを使う
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println()
}
