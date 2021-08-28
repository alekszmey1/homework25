package main

import (
	"fmt"
)

func main()  {
	a:=4
	aptr := returnIntByPointer(a)
	fmt.Println(aptr)//печатает адрес новой переменной
	fmt.Println(*aptr)//печатает значение переменной v
	}

func returnIntByPointer(a int)*int  {
	v:=a //присваивает переменной v значение 4
	return &v//возвращают адрес новой переменной
}