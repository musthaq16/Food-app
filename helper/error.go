package helper

import "fmt"

func ErrorPanic(err error) {
	if err != nil {
		fmt.Println("there is error and panicing....", err)
	}
}
