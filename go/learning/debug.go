package main

import (
	"fmt"
	"runtime/pprof"
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

func changeString(i int, s []string, done chan string) {
	//for _, p := range pprof.Profiles() {
	//	fmt.Println(p.Name())
	//	fmt.Println(p.Count())
	//}

	p := pprof.Lookup("goroutine")
	fmt.Println(p.Count())
	s[i] = "new"
	done <- "done!"
}

func main() {

	var s = []string{"1", "2"}
	fmt.Println(s)
	done := make(chan string)

	for i := 0; i < len(s); i++ {
		go changeString(i, s, done)
	}

	for i := 0; i < len(s); i++ {
		<-done
	}

	fmt.Println(s)

}
