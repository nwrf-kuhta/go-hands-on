package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

// SetValue is Mydata method.
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// PrintData is Mydata method.
func (md *Mydata) PrintData() {
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	} else {
		fmt.Println("** This is Nil value.**")
	}
}

// nilレシーバについて
func main() {
	var ob *Mydata
	ob.PrintData()
	ob = &Mydata{}
	ob.SetValue(map[string]string{
		"name": "Jiro",
		"data": "123 456 789",
	})
	ob.PrintData()
}
