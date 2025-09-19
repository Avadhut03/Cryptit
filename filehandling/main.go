package main

import (
	"os"
)

func main() {
	// path := filepath.Join("dir1","dir2","text.txt") 

	// println(filepath.Dir(path))
	// println(filepath.Base(path))
	// println(filepath.IsAbs(path))
	// println(filepath.IsAbs("/dir/file"))
	// println(filepath.Ext(path))

	// fileinfo , err:= os.Stat("C:/Users/HP/OneDrive/Desktop/Cryptit/test.txt")
	// if err!= nil {
	// 	println("Error:", err.Error())
	// 	return
	// }
	// println(fileinfo.Name())
	// println(fileinfo.Size())
	// println(fileinfo.IsDir())
	// println(fileinfo.Mode())
	
	// println(path)
	path:= "C:/Users/HP/OneDrive/Desktop/Cryptit/test.txt"
	// data,err:= os.ReadFile(path)
	data, err:= os.Open(path)
	if err!= nil {
			println("Error:", err.Error())
			return
		}
	println((data))
	b:= make ([]byte,4)
	for{
		n,err:=data.Read(b)
		if err!= nil {
			println("Error:", err.Error())
			return
		}
		println(string(b[0:n]))
	}
	
}