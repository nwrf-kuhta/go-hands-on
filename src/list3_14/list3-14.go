package main

import "fmt"

// Data is interface.
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// Mydata is Struct.
type Mydata struct {
	Name string
	Data []int
}

// Initial is init method.
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Check is method.
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]", md.Name)
}

// インターフェースを利用する
func main() {
	var ob = Mydata{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
	ob.Check()
	// MydataをDataとして扱う
	var ob2 Data = new(Mydata)
	ob2.Initial("Sachio", []int{11, 22, 33})
	ob2.PrintData()
	// ob2.Check() : ob2は、Data型（インターフェース）であるため、Mydata固有のメソッドを使うことはできない
}
