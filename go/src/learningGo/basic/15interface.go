package main

import (
	"math"
	"fmt"
)

//interface is collection of method signature
type geometry interface {
	//list of method signatures
	area() float32
	perimeter() float32
}

type rect struct {
	width, length float32
}

type circle struct {
	radius float32
}

//now implement area and perimeter
func (r rect) area() float32 { // it simply means that area function has been inserted inside class rect, or area function
    //has been registered with rect and r.area() can be called
	return r.length * r.width
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

//implemented perimiter function
func (r rect) perimeter() float32  {
	return 2*(r.width+r.length)
}
func (c circle) perimeter() float32 {
	return 2*math.Pi*c.radius
}

func main()  {
	r := rect{10,15}
	c := circle{10}
	fmt.Println(r.area())
	fmt.Println(c.area())

	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())

	r1 := rect{width:12,length:14}
	c1 := circle{radius:10}
	//since r and circle both implements methods of rectangle and circle, both can be passed as geometry (typecast!!)
	//if rectangle or circle doesnt implement any one method then it can't be passed as geometry type
	measure(r1)
	measure(c1)
}

func measure(g geometry)  {
	fmt.Println(g.perimeter())
	fmt.Println(g.area())
}

