package main

import "fmt"


type Shape interface{
	getArea() float64
}

type Triangle struct{
	height float64
	base float64
}

type Square struct{
	sideLength float64
}

func main(){
	fmt.Println("Shapes assignment!!")
	sq := Square{sideLength: 2}
	tr := Triangle{height: 2, base: 2}
	printArea(sq)
	printArea(tr)
}


func (s Square) getArea() float64{
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64{
	return 0.5 * t.base * t.height
}

func printArea(sh Shape){
	fmt.Println(sh.getArea())
}