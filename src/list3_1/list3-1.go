package main

import "fmt"

// ポインタを使ってみる
func main() {
	n := 123
	// 変数の頭に&をつけることでポインタを取得
	p := &n
	fmt.Println("number: ", n)
	fmt.Println("pointer: ", p)
	// ポインタ変数の前に*をつけるとそのポインタにある値を取得
	fmt.Println("value:", *p)
	// ポインタは演算したりはできない。
	// ポインタは「変数の値があるアドレスを示す値」である。
	// ポインタは方が違っても全てint値
}
