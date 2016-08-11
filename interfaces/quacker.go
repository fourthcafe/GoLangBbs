package interfaces

import "fmt"

func MainQuacker() {
	p := Person{
		name: "ko",
		age:  32,
	}
	q := Duck{
		nick: "Donald",
		area: "서울",
	}

	if v, ok := interface{}(p).(functional); ok {
		fmt.Printf("v: %+v\n", v)
		introduce(v)
	}

	introduce(p)
	introduce(q)
}

type functional interface {
	info()
	feature()
}

func introduce(q functional) {
	q.info()
	q.feature()
}

type Person struct {
	name string
	age  int
}

func (this Person) info() {
	fmt.Println(this.name, this.age)
}

func (this Person) feature() {
	fmt.Println("진상")
}

type Duck struct {
	nick string
	area string
}

func (this Duck) info() {
	fmt.Println(this.nick, this.area)
}

func (this Duck) feature() {
	fmt.Println("날 수 있음")
}
