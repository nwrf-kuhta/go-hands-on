package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// レンジについて
func main() {
	x := input("input data: ")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
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
