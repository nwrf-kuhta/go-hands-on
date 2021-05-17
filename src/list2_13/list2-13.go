package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ラベルとgoto
func main() {
	t := 0
	x := input("type a number: ")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err
	}
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
	return

err:
	fmt.Println("ERROR!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg)
	scanner.Scan()
	return scanner.Text()
}
