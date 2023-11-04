package main

import "fmt"

type number int

type car struct {
	carName  string
	carModel string
	year     int
}

func (num number) isEven() bool {
	return num%2 == 0
}

func (v number) isOdd() bool {
	return v%2 != 0
}
func(c car) carDetails() string{
	return c.carName+"  "+ c.carModel
}

func Oops() {
	fmt.Println("OOps starts here")

	var num number
	var v number
	var c car
	c=car{
     carName:"TAta",
	 carModel:"Petrol",
	 year:2023,

	}
	num = 10
	fmt.Println(num)
	fmt.Println(num.isEven())
	fmt.Println(v.isOdd())
	fmt.Println(c.carDetails())

}
