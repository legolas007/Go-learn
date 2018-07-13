package main

import (
	"fmt"
)

type test struct {
	name string
	age int
	Contact struct{
		phone ,city string
	}
	hus
}
type hus struct {
	qps string
}
func main()  {
	b := &struct {
		name string
		age int
	}{
		name:"df",
		age:56,
	}
	a := test{
		name:"uh",
		age:23,
		hus:hus{"ssd"},
	}
	a.Contact.phone = "2323"
	a.Contact.city="ws"
	fmt.Println(a)
	fmt.Println(*b)

	A(1,2,3)
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
func A(a ...int)  {
	fmt.Println(a)
	fmt.Println(a[0])
}
func closure(x int) func(int) int  {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n",&x)
		return x+y
	}
}
