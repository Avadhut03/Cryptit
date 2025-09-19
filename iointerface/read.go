package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// r := strings.NewReader("Hello, Reader!")
	// buf:= make([] byte,4)
	// for{
	// n,err:= r.Read(buf)
	// fmt.Println(string(buf[:n]),err)
	// if err!=nil{
	// 	break
	// }
	// }
	f, _ := os.Create("test.txt")
	defer f.Close()
	f.Write([]byte("Hello World"))
	r:= strings.NewReader("Some io.reader stream to be read")
	if _, err:= io.Copy(os.Stdout,r); err!=nil{
		log.Fatal(err)
	}
}