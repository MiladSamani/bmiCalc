package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Title = "BMI Calculator"
const Line = "--------------"
const WeightDescription = "Please Enter Your Weight (kg) : "
const heightDescription = "Please Enter Your Height (m) : "

var reader = bufio.NewReader(os.Stdin)

func main()  {
	fmt.Println(Title)
	fmt.Println(Line)
    fmt.Print(WeightDescription)
	weightInput , _ := reader.ReadString('\n')
	fmt.Print(heightDescription)
	heightInput , _ := reader.ReadString('\n')
	fmt.Print(weightInput)
	fmt.Print(heightInput)
	weightInput = strings.Replace(weightInput , "\n" ,"" , -1)
	heightInput = strings.Replace(heightInput , "\n" ,"" , -1)
	weight , _ := strconv.ParseFloat(weightInput , 64)
	height , _ := strconv.ParseFloat(heightInput , 64)
	bmi:= weight / (height * height)
	fmt.Printf("Your BMI : %.2f" , bmi)
}
