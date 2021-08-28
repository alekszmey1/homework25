package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main()  {
	var name string
	var age int

	flag.IntVar(&age, "age", 0 , "set age")
	flag.StringVar(&name, "name", "default name", "set name")

	flag.Parse()
	fmt.Println(name,age)

	printByPointer(&name,&age)

}

func printByPointer(n*string, a*int)  {
	if n == nil || a==nil {
		return
	}
	fmt.Println("hello. my name is "+*n+ " and age is "+strconv.Itoa(*a))
}