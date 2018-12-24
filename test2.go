package main

import (
	"fmt"

)
const (
	b float64 = 1 << (iota * 10)
	a
	MB
	GB
	)
func main(){
	a :=1
	var p *int = &a
	fmt.Println(p)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(MB)
	fmt.Println(GB)
}
