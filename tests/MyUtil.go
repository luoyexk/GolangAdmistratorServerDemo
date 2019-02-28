package test

import "fmt"

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:[ ", err, " ]")
		panic(err)
	}
}

func GetTime() {

}
