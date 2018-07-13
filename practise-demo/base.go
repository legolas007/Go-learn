package main

import (
	"fmt"

)

func main()  {
	var b [2]int
	c := [...]int{0:1,1:2,2:3}
	fmt.Println(c)
	fmt.Println(b)
	a := [...]int{99:1}
	var p *[100]int  = &a
	fmt.Println(*p)
	fmt.Println(a)
	pp := new([10] int)
	pp[1] = 2
	fmt.Println(*pp)
	aa := [10]int{}
	aa[1] = 2
	fmt.Println(aa)
	d := [2][3]int{
		{1,1,1},
		{},
	}
	fmt.Println(d)
	s := [10]int{9:1}
	s1 := s[:]
	fmt.Println(s1)
	s2 := make([]int,3,10)
	fmt.Println(len(s2),cap(s2))
	s2 = append(s2,23,23 )
	copy(s1,s2)
	fmt.Println(s1)
	fmt.Println(s2)
	m := make(map[int]string)
	m[1] = "ok"
	fmt.Println(m)
	delete(m,1)
	var mm map[int]map[int]string
	mm = make(map[int]map[int]string)
	m1,ok := mm[2][1]
	if !ok{
		mm[2] = make(map[int]string)
	}
	mm[2][1] = "good"
	m1,ok = mm[2][1]
	fmt.Println(m1,ok)
	sm := make([]map[int]string,5)
	for _,v := range sm{
		v = make(map[int]string,1)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)

}