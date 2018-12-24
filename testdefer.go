package main

import "fmt"

func main(){
	/*
	for i :=0 ;i<3 ;i++ {
		defer fmt.Println(i)

		/*defer func(){
			fmt.Println(i)
		}
	}*/
	A()
	B()
	C()
}
func A() {
	fmt.Println("func a")
}
func B() {
	defer func(){
		if err :=recover() ;err!=nil {
			fmt.Println(err)
			}
		}()
	panic("panic in b")
}
func C() {
	fmt.Println("cccc")
}
func tt(){
	var fs =[4]func(){}
	for i:=0;i<4;i++{
		defer fmt.Println("defer i=" ,i)
		defer func() {
			fmt.Println("defer_closure i =",i)
		}()
		fs[i] = func(){
			fmt.Println("closeure i=",i)
		}
	}
	for _, f :=range fs{
		f()
	}
}
