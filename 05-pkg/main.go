package main

import (
	"fmt"
	"05-pkg/lib"
)
//import mylib "05-pkg/lib"
//will give an alias to the package at 05-pkg/lib

//ignore importing package by
//import _ "05-pkg/lib"

func main(){
	fmt.Println(lib.IsDigit('1'))
	fmt.Println(lib.IsDigit('a'))
}
