package helpers

import "fmt"

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}