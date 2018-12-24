package main

import "fmt"

func main(){
	//var m map[int]string
	m := make(map[int]string)
	m[1] = "ok"
	fmt.Println(m)
	m1 := make(map[int]map[int]string)
	a, ok :=m1[2][1]
	if !ok {
		m1[2] = make(map[int]string)
	}
	m1[2][1] = "lalalla"
	a, ok =m1[2][1]
	fmt.Println(a,ok)
	fmt.Println(m1,a)
}
