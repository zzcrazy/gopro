package main

import "fmt"

func main(){
	h :=[]int{1,3}
	c := B
	c(h)
	fmt.Println(h)
	A(1,2)
	//匿名函数
	w :=func(){
		fmt.Println("hell")
	}
	w()
}
func A(a ...int)  {
	fmt.Println(a)
	f :=closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
func B(s []int)  {
	s[0] = 4
	s[1] = 5
	fmt.Println(s)
}
func closure(x int) func(int) int{
	return func(y int) int{
		return x+y
	}
}


