//encrypt package consist of all the algorithms used for decryption
package encrypt
//encrypt by reducing ascii value by 4

func Nimbus(str string) string{
	encrypted:=""
	for _,c:=range str{
		ascii:=int(c)
		ch:=string(ascii+4)
		encrypted +=ch
	}
	return encrypted
}