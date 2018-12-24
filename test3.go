package main

import "fmt"

func main(){
	a := 1
	if a :=2; a>1 {
		fmt.Println(a)
	}
	fmt.Println(a)
	b := 1
	for {
		b++
		if b>3{

			fmt.Println(b)
			break
		}
	}
	fmt.Println("________________---------")
	for i :=0 ; i < 3; i++{
		fmt.Println(i)
	}
	fmt.Println("**************--------")
	f := "papapa"
	l := len(f)
	for i :=0 ; i < l; i++{
		fmt.Println(f[i])
	}
}
