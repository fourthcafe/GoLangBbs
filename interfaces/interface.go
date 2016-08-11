package interfaces

import "fmt"

func Main() {
	var i myInt = 5
	r := rectangel{10, 20}

	var print Printer

	print = i
	print.print()

	print = r
	print.print()

	// 선언하면서 초기화
	iP := Printer(i)
	iP.print()

	// 슬라이드 형태
	var slideInt myInt = 98
	slideRect := rectangel{20, 100}
	pA := []Printer{slideInt, slideRect}
	for idx, _ := range pA {
		pA[idx].print()
	}

	for _, value := range pA {
		value.print()
	}
}

type myInt int

func (this myInt) print() {
	fmt.Println(this)
}

type rectangel struct {
	width, height int
}

func (this rectangel) print() {
	fmt.Println(this.width, this.height)
}

type Printer interface {
	print()
}
