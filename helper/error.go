package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("masuk3")
		panic(err)
	}
}
