package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// fotの初期化と後処理
func main() {
	x := input("type a number: ")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print("1から" + x + "の合計は、")
	} else {
		return
	}
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println(t, "です。")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}
