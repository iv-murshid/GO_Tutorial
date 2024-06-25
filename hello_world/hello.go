package hello_world

import "fmt"

func Hello(msg string) (string,error) {
	fmt.Println(msg)
	return msg,nil
}