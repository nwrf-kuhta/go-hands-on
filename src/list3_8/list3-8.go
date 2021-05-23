package main

import "fmt"

// MyData is structure.
type Mydata struct {
	Name string
	Data []int
}

// 構造体と参照渡し
func main() {
	taro := Mydata{
		"Taro",
		[]int{10, 20, 30},
	}
	fmt.Println(taro)
	taro = rev(taro)
	fmt.Println(taro)
	rev2(&taro)
	fmt.Println(taro)
}

func rev(md Mydata) Mydata {
	od := md.Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return md
}

func rev2(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
