package main

import (
	"fmt"
	"series"
)

func main() {
	//fmt.Println(os.Args)
	fb := series.GetFibonacciSeries(10)
	fmt.Println(fb)
}
