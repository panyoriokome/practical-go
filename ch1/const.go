package main

import "fmt"

func main() {
	var a = 1       // イコールの右側の1が定数
	var b = 1 + 100 // 定数の中で演算子を使った計算も可能

	var c = 1 // 定数のデフォルトの型になる
	var d int32 = 100

	fmt.Println(a, b, c, d)

	// 型付きconst変数を定義する
	type ErrorCode int

	const (
		f    int       = 10
		code ErrorCode = 10
	)

	// constで定義すると値の書き換えはできなくなる
	// f = 2

	//
	const (
		g int32          = 4294967295
		h []int          = []int{1, 2, 3}
		i map[string]int = map[string]int{
			"Tokyo": 10,
		}
		j = sampleFunction()
	)
}

func sampleFunction() int {
	return 2
}
