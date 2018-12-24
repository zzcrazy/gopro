package main

import (
	"fmt"
	"math"
)
type (
	byte int8
	rune int32
	文本 string
)
func main(){
	var a bool
	var b [2]byte
	fmt.Println("hello" )
	fmt.Println( a )
	fmt.Println( math.MaxInt32)
	fmt.Println( b )
	fmt.Println( "---------------------" )
	var c 文本
	c = "你好"
	fmt.Println( c )
	fmt.Println( "---------------------" )
	var aa float32 =100.1
	fmt.Println(int(aa))
	//fmt.Println(bool(aa))
}