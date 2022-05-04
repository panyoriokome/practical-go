package main

import "fmt"

func main() {
	example1()
	exmaple2()
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
