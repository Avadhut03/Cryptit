package main

import (
	"fmt"
	"strings"
)

func main(){
	learning := "learning standard library in go "
	fun:= "library in go"
	result:= strings.Contains(learning, fun)
	fmt.Println(result)
	// fmt.Println(strings.Count(learning, "go") )
	ans:= strings.ReplaceAll(learning,fun,"library in python")
	fmt.Println(ans)



		


}