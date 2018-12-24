package main

import "fmt"

func main(){
	//var a [2]int
	a := [3]int{1,2:2}
	var p *[3]int = &a
	var b [1]int
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(p)
	d := new([10]int)
	fmt.Println(d)
	c :=[5]string{}
	//c[1]="asdf"
	fmt.Println(c)
	e :=[2][3]int{
		{1,23,2},
	}
	fmt.Println(e)
	}
