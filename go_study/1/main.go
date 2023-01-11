package main

import "fmt"


type MF float64
type cat struct{
	name string
	age int
}

func main() {
	newCat := cat{name: "hello",age: 1}
	fmt.Println(newCat)
	var p MF = 1.4
	p.show()
}

func point(n *int) {
	*n = *n + 1
}
func other(){
	fmt.Println("vim-go")

	type MF float64
	var b MF
	var n int
	point(&n)
	fmt.Println(n)
	b=1.1
	fmt.Println(b)
}
func (u MF) show(){
	fmt.Println(u)
}
