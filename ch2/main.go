package main

import "fmt"

func main() {
	type SourceType int

	var int int = 1
	var source1 SourceType = 1
	// 型が違うからintの変数を代入できない
	// cannot use int (variable of type int) as SourceType value in variable declaration
	// var source2 SourceType = int

	fmt.Println(int, source1)
}
