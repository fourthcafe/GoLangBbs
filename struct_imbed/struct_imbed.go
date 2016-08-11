package struct_imbed

import "fmt"

func Main() {
	stuP := NewStudentIsA()
	stuP.introStd()
}

// Person 을 is-a 관계로
type StudentIsA struct {
	Person
	school string
	grade  int
}

// Person 을 has-a 관계로
type StudentHasA struct {
	p      Person
	school string
	grade  int
}

type Person struct {
	name string
	age  int
}

func (this StudentIsA) introStd() {
	fmt.Println("school: ", this.school)
	fmt.Println("grade: ", this.grade)
	this.greeting()
}

func (p Person) greeting() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
}

func NewStudentIsA() *StudentIsA {
	return &StudentIsA{
		Person: Person{
			name: "ko",
			age:  32,
		},
		school: "Is A",
		grade:  8,
	}
}

func NewStudentHasA() *StudentHasA {
	return &StudentHasA{
		p: Person{
			name: "kwak",
			age:  27,
		},
		school: "Has A",
		grade:  8,
	}
}
