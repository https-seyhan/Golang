package main

import (
	"fmt"
	"errors"
	"math"
)
type person struct {
	name string
	age int
}

func main() {

	p :=person {name: "Saul", age: 35}
	fmt.Println(p)
	
	//fmt.Println("Test")
	x := 5
	//var y int = 7
	//var sum int = x + y
	
	//fmt.Println(sum)
	
	if x > 7 {
		fmt.Println(x)
	}
	
	var a [5]int
	a[1] = 7
	b :=[]int{1, 2, 3, 4, 5}
	b = append(b, 13)
	fmt.Println(a)
	fmt.Println(b)
	
	vertices := make(map[string]int)
	
	vertices["triangle"] =2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	delete(vertices, "square")
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	
	for i:=0; i <5; i++{
		fmt.Println(i)
	}
	
	m:=make(map[string]string)
	m["a"] = "alpha"
	m["b"] ="beta"
	
	for key, value := range m{
		fmt.Println(key, value)
	} 
	
	result :=sum(4,5)
	fmt.Println(result)
}

func sum(x int, y int) int {
	
	
	result, err := sqrt(16)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative Number")
	}
	
	return math.Sqrt(x), nil
}

