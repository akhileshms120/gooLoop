package main
import "fmt"
func Branching(){

	//if statement
   i:=6

   if i>6{  //condition failed, here

	fmt.Println("i>6")
   } else if i<6{
	
	fmt.Println("i<6")
   }else{
    fmt.Print("failed")
   }
   fmt.Println()
  //switch'

  switch i{
  case 3:
	fmt.Println("Print 3")
  case 6:
	fmt.Print("Print 6")
  default:
	fmt.Println("Print default")
  }
  fmt.Println()
  
  switch {
case i<3:
  fmt.Println("Print 3")
case i<6:
  fmt.Print("Print 6")
default:
  fmt.Println("Print default")
}

switch {
case i<3,i<7:
  fmt.Println("Print 3")
case i>6:
  fmt.Print("Print 6")
default:
  fmt.Println("Print default")
}

///goTo statement


if i>5 {
	fmt.Print("i greater than 7")
	goto mylabel
}
//m:=54  if i decelare varaible here goto not work
mylabel:
  if i<7{
	println("after label c:7")
  }
}