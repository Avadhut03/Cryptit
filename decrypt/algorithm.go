package decrypt
func Nimbus(str string) string{
	decrypted:=""
	for _,c:= range str{
		asc:=int(c)
		ch:= string(asc-4)
		decrypted+=ch
	}
	return decrypted
}