package main

type Shape interface{
	Area() float64
	Permiter() float64
}
type Rectangle struct
{
	length, width float64
}
func (r Rectangle) Area() float64{
	return r.length*r.width
}
func (r Rectangle) Permiter() float64{
	return 2*(r.length+r.width)
}

func simple(){
	var s Shape = Rectangle{length:10, width:5}
	println("Area:", s.Area())
	println("Permiter:", s.Permiter())

}