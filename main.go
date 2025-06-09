package main

import "fmt"

func intersection(slicea []int, sliceb []int) []int {
	var m []int
	for _, i := range slicea {
		for _, i2 := range sliceb {
			if i2 == i {
				m = append(m, i)
			}
		}
	}
	return m
}

func main() {
	slicea := []int{5, 1, 2, 5}
	var n = 3
	slicea = append(slicea[:n], slicea[n+1:]...)
	sliceb := []int{4, 2, 5, 1, 1, 2}
	var y = 1
	var x = 3
	sliceb = append(sliceb[:y], sliceb[y+1:]...)
	sliceb = append(sliceb[:x], sliceb[x+1:]...)
	fmt.Println(slicea, sliceb)
	sliceq := append(slicea, sliceb...)
	w := sliceq[0:4]
	fmt.Println(intersection(slicea, sliceb))
	fmt.Println(w)
}
