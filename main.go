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
	for i := 0; i < len(a); i++ {
		if a[i] != b[0]{
			fmt.Println("руна ", a[i], "не равна", b[0])// для отслеживания выполнения задачи
			continue
		} else {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					fmt.Println("найдено следующее совпадение,руна ", a[i], " равна", b[j])//для отслеживания выполнения задачи
					i++
					if j == len(b)-1 {
						fmt.Println("найдено 100 % совпадение")
						break
					}
					continue
				} else {
					i = i - j + 1
					j = 0
					return
				}
			}
		}
	}
}

