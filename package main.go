package main

import "fmt"

func main() {
	var i1 int = 42
	var f1 float64 = 13.37
	var s1 string = "Stroka"
	var b1 bool = true

	var i2 int
	var f2 float64
	var s2 string
	var b2 bool

	i3 := 5
	f3 := 5.2
	s3 := "stroka"
	b3 := false

	fmt.Printf("%v - %T\n", i1, i1)
	fmt.Printf("%v - %T\n", f1, f1)
	fmt.Printf("%v - %T\n", s1, s1)
	fmt.Printf("%v - %T\n", b1, b1)

	fmt.Printf("%v - %T\n", i2, i2)
	fmt.Printf("%v - %T\n", f2, f2)
	fmt.Printf("%v - %T\n", s2, s2)
	fmt.Printf("%v - %T\n", b2, b2)

    fmt.Printf("%v - %T\n", i3, i3)
	fmt.Printf("%v - %T\n", f3, f3)
	fmt.Printf("%v - %T\n", s3, s3)
	fmt.Printf("%v - %T\n", b3, b3)

a , b := 1,2
tmp := a
a = b
b =tmp
a , b=b , a
fmt.Printf("% - %T\n" , a , a)
fmt.Printf("% - %T\n" , b , b)
 

}