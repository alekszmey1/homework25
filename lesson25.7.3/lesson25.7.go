package main

import "fmt"

func main() {
	integer := 10
	fmt.Println(integer)
	integerPointer := &integer //присваивается адрес хранения переменной integer
	fmt.Println(integerPointer)
	integerFromPointer := *integerPointer //показывается переменная, которая хранится в другой переменно
	fmt.Println(integerFromPointer)
	v := 0
	fmt.Println(v)
	changeValueByValue(v)
	fmt.Println(v)
	changeValueByPointer(&v)
	fmt.Println(v)
}
	func changeValueByValue(v int){//не поменяет значение т.к. нет ссылки на переменную
		v=10
	}

	func changeValueByPointer(v *int){//поменяет значение внутри переменной v
		*v = 10
	}

