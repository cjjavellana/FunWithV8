package main

import (
	"log"
	v8 "rogchap.com/v8go"
)

func SumWithGenerics[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)
	_, err := ctx.RunScript("const multiply = (a, b) => a * b", "math.js")
	if err != nil {
		log.Panic(err)
	}

	r, err := ctx.RunScript("multiply(3.14159, 4).toFixed(5)", "main.js")
	if err != nil {
		log.Panic(err)
	}
	println(r.DetailString())

	var m = make(map[string]float64)
	m["s"] = 1.2
	m["t"] = 1.5
	println(SumWithGenerics(m))

	var n = make(map[string]int64)
	n["s"] = 100
	n["t"] = 200
	println(SumWithGenerics(n))
}
