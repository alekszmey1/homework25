package main

import (
	"flag"
	"fmt"
)

/*Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:
 go run main.go --str "строка для поиска" --substr "поиска"
Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.*/

func main()  {
	var text string
	var partText string

	flag.StringVar(&text, "str", "default name", "set name")
	flag.StringVar(&partText, "substr", "default name", "set name")

	flag.Parse()
	fmt.Println(text,partText)

	strrune2 :=[]rune(partText)
	strrune := []rune(text)
	coincidence(strrune,strrune2)
}

func coincidence(a[]rune, b[]rune) {
	for i := 0; i < len(b); i++ {
		for j := 0; j <len(a); j++ {
			if a[j] == b[i]{
				fmt.Println("найдено совпадение. буква", string(a[i]))
			}
		}
	}
}

