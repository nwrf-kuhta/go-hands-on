package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ショートステートメント付き「if」
func main() {
	x := input("type a number: ")
	if n, err := strconv.Atoi(x); err == nil {
		fmt.Print(x + "は、")
		if n%2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("ERROR!!")
	}
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}
