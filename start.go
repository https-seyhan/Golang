package main

import (
	"fmt"
)

func main() {

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
}
