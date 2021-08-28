package main

import "fmt"

//шифр цезаря
func main ()  {
	message := "sometext"
	for i := 0; i < len(message); i++ {
		c:= message[i]
		if c >= 'a' && c <='z'{
			c = c+13
			if c > 'z'{
					c= c-26
			}
			fmt.Printf("%+c",c)
		}

	}

}
