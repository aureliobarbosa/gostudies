package main

import "fmt"

func main(){
	// testing attribution to slices in different ways
	type MyInt []int64

	var v0 MyInt
	v0 = MyInt{1,2,8,4}	
	fmt.Println(v0)

	var v1 MyInt = MyInt{1,72,38,14}	
	fmt.Println(v1)

	v2 := MyInt{7,3,0,12,78}
	fmt.Println(v2)

	var v3 MyInt
	v3 = append(v3, 1,2,3,4,5,6,7,8,9)
	fmt.Println(v3)	
}