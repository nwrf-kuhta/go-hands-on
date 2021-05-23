package main

import "fmt"

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

// newとmake
func main() {
	// newで構造体を作成する
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	// makeで配列・スライス・マップ・チャンネルなどの初期化を行う
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)
}
