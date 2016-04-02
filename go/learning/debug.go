package main

import (
	"fmt"
	"runtime"
)

//var g byte = '1'
//
//type rect struct {
//	width, height int
//}
//
//func (r *rect) area() int {
//	return r.width * r.height
//}
//
//func (r *rect) perim() int {
//	return 2 * (r.height + r.width)
//}
//
//func (r rect) change() {
//	r.width = 11
//}
//
//func (r *rect) changeP() {
//	r.width = 12
//}

func main() {


	cpus := runtime.NumCPU()
	fmt.Println(cpus)


}
