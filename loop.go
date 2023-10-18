

package main
import "fmt"

func loop() {
    fmt.Println("Loops Demo")

    // Infinite loop
    for {
        fmt.Print("Infinite for loop")
        // For exit, use break
        break
    }

    // Condition-based loop
    var num int
    num = 1

    for num <= 5 {
        fmt.Print(num) // Use fmt.Print instead of fmt.Println
        num++
    }

    // Add a newline character to separate the numbers
    fmt.Println()


	//counter based looop
	
	for i:=0;i<=5;i++{
		fmt.Println(i)
	}

	//collection based loop

	numbers:=[4]int{
		0,100,200,300,
	}

	for _,v:=range numbers{      //for i,v:=range numbers{   
		//fmt.Println("index",i,"Value",v)
		fmt.Println("Value",v)

		//for unused variable you can declare _(underscore)
	}

	//loop maps

	letters:= map[string]string{
      "a":"A",
	  "b":"B",
	  "c":"C",
	}

	for k,v:=range letters{
		fmt.Println(k,v)   //skip K use underscore
	}
}
