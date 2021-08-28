package main

import (
	"fmt"
	"strconv"
)

func main()  {
	lange := []string{"python","go", "c", "js"}
	compiledLange := lange [1:3]//выделение подмножества из среза
	fmt.Println(lange)//срез
	fmt.Println(compiledLange)//подмножество
	students:= []string{"mary", "cary","sonya", "petro", "vedro","mark"}//композитный литерал
	//students2 := make([]string,0,0)//создание слайса через функцию make
	studentsSubslice := students[3:]
	fmt.Println(studentsSubslice)
	studentsSubslice[1]="aleks"
	fmt.Println(students, studentsSubslice)
	var b byte = 255
	fmt.Printf("%b\n",b)
	var c byte = 65
	fmt.Printf("%b\n",c)
	var d byte = 2
	fmt.Printf("%b\n",d)
	hello :="hello"
	helloBiteSlice:=[]byte(hello)
	fmt.Println(helloBiteSlice)
	por := 123456
	por2 := strconv.Itoa(por)
	fmt.Println(por2)
	por3 := []byte(por2)
	fmt.Println(por3)
	for i := 0; i < len(por3); i++ {
		var pol byte = por3[i]
		fmt.Printf("%b\n",pol)
	}
	hi:="привет"
	por4 := []byte(hi)
	for i := 0; i < len(por4); i++ {
		var pol byte = por4[i]
		fmt.Printf("%b\n",pol)
	}
	const lol = "🤣"//руна

	fmt.Printf("%s\n",lol)//запринтили как строку
	fmt.Printf("%+q\n",lol)// запринтили как кодовую точку юникода
	for i := 0; i < len(lol); i++ {
		fmt.Printf("%x\n",lol[i])//шестнадцатиречиное представление байт из которых состоит смайлик
		fmt.Printf("%b\n",lol[i])

	}
	fmt.Printf("\n")
	character := 'ф'
	fmt.Println(character)
	fmt.Println(string(character))
	str := "привет"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d ",str[i])//печатаем значение байтов
		}
	for i := 0; i < len(str); i++ {
		fmt.Printf("%b ", str[i])// печатаем байты побитого
	}
	strrune := []rune(str)
	for _, r := range strrune {
		fmt.Println(r, string(r))//присвоенный номер в юникоду utf 8 и сами руны
	}
}
