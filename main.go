package main

import (
	"flag"
	"fmt"
	"strings"
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
	worlds := strings.Fields(text)
	for _, world := range worlds {
		strrune := []rune(world)
		coincidence(strrune,strrune2)
	}
}
func coincidence(a[]rune, b[]rune) {
	if len(a) == len(b) {
			for i := 0; i < len(a); i++ {
			if a[i] != b[i]{
				break
			} else if i == len(a)-1{
				fmt.Println("найдено совпадение")
			}
		}
	}
}

