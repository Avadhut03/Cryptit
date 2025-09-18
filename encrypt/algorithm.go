package encrypt

func Nimbus(str string) string{
	encrypted:=""
	for _,c:=range str{
		ascii:=int(c)
		ch:=string(ascii+4)
		encrypted +=ch
	}
	return encrypted
}