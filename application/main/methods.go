package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	r1 := rect{width: 10, height: 5}
	r2 := rect{width: 10, height: 5}
	r1.modifyNormal()
	r2.modifyPointer()
	fmt.Println(r1)
	fmt.Println(r2)

	r3 := rect{width: 10, height: 5}
	r4 := rect{width: 10, height: 5}
	modifyNormal(r3)
	modifyPointer(&r4)
	fmt.Println(r3)
	fmt.Println(r4)
}

func (r rect) modifyNormal() {
	r.width += 10
	r.height += 5
}
func (r *rect) modifyPointer() {
	r.width += 11
	r.height += 6
}

func modifyNormal(r rect) {
	r.width += 10
	r.height += 5
}
func modifyPointer(r *rect) {
	r.width += 11
	r.height += 6
}
