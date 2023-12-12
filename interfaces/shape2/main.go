package main

import "fmt"

type Shape interface{
	Area() float64
}

type Square struct{
	side float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length float64;
	breadth float64;
}


func (s Square) Area() float64{
	return (s.side*s.side)
}

func (c Circle)Area() float64{
	return 3.14*c.radius*c.radius	
}


func (r *Rectangle)Area() float64{
	return r.length * r.breadth
}

func printArea(s Shape){
	fmt.Println(s.Area())
}

func main(){
	s := Square{side: 5}
	printArea(s)
	c := Circle{radius: 5}
	printArea(c)

	r := Rectangle{length:5, breadth:5}
	printArea(&r)
}