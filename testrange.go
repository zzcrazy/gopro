package main

import (
	"fmt"
	//"sort"
)

func main(){
	//sm :=make([]map[int]string,5)
	/*
	for _,v :=range sm{
		v =make(map[int]string,1)
		v[1] ="hahah"
		fmt.Println(v)
	}

	for i :=range sm{
		sm[i] =make(map[int]string,1)
		sm[i][1] ="hahah"
		fmt.Println(sm[i])
	}*/
	m := map[int]string{1:"a",2:"b",3:"c"}
	//s :=make([]int,len(m))
	s1 :=make(map[string]int)
	for k,v :=range m{
		s1[v]=k
	}
	fmt.Println(s1)
/*
	i :=0
	j :=0
	for k,_ :=range m{
		s[i] = k
		i++
	}
	for k,_ :=range m{
		fmt.Println(m[k])
		s1[m[k]] = k
		j++
	}
	//sort.IntsAreSorted(s)
	fmt.Println(s)
	fmt.Println(s1)*/
	}
