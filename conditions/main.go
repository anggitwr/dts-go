package main

import "fmt"

func main(){
	var curentYear = 2021

	if age := curentYear - 1998; age < 17{
		fmt.Printf("kamu belum punya sim")
	} else {
		fmt.Printf("kamu sudah punya sim")
	}
}