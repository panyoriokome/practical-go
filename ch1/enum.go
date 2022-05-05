package main

import "fmt"

func main() {
	example1()
	exmaple2()
	example3()
}

func example1() {
	fmt.Println("exmaple1:")

	type CarType int

	const (
		Sedan CarType = iota + 1
		Hatchback
		MPV
		SUV
		Crossover
		Coupe
		Convertible
	)

	var t CarType
	t = SUV
	fmt.Println(t) // 4
}

func exmaple2() {
	fmt.Println("exmaple2:")

	const (
		a = iota
		b
		c
		_ // 空行はスキップされるが、その分数値がインクリメントされる
		d
		e = iota
	)

	fmt.Println(a, b, c, d, e)

	const (
		f = iota // 新しいブロックだとまた0からカウントされる
		g
		h
	)

	fmt.Println(f, g, h)
}

func example3() {
	fmt.Println("example3")
	type CarOption uint64

	const (
		GPS          CarOption = 1 << iota
		AWD                    // 2
		SunRoof                // 4
		HeatedSeat             // 8
		DriverAssist           // 16
	)

	var o CarOption
	o = SunRoof | HeatedSeat
	if o&SunRoof != 0 {
		fmt.Println("サンルーフ付き")
	}

	fmt.Println(AWD, SunRoof, HeatedSeat, DriverAssist)
}
