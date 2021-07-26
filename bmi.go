package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main()  {
	fmt.Println("BMI CALCULATOR")
	fmt.Println("--------------")
    fmt.Print("Please Enter Your Weight (kg) : ")
	weightInput , _ := reader.ReadString('\n')
	fmt.Print("Please Enter Your Height (m) : ")
	heightInput , _ := reader.ReadString('\n')
	fmt.Println(weightInput)
	fmt.Println(heightInput)
}
