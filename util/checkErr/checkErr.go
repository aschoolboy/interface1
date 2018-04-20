package checkErr

import "fmt"

func CheckErr(err error, s int) {
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
}