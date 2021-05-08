package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 条件のないswitch
func main() {
	x := input("type a number: ")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "は、")
	} else {
		fmt.Println("ERROR!!")
		return
	}
	switch {
	case n%2 == 0:
		fmt.Println("偶数です。")
	case n%2 == 1:
		fmt.Println("奇数です。")
	}
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}
