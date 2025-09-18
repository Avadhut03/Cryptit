package main

import (
	"fmt"

	"github.com/Avadhut03/Cryptit/decrypt"
	"github.com/Avadhut03/Cryptit/encrypt"
)

func main(){
	encrypted:=encrypt.Nimbus("Avadhut")
	fmt.Println("Encrypted String:",encrypted)
	fmt.Println(decrypt.Nimbus(encrypted))

}