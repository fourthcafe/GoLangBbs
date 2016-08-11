package closure

import "fmt"

func Main() {
	f := calc(2, 4)
	result := f(9)
	fmt.Println(result)
}

func calc(a, b int) func(k int) int {
	return func(k int) int {
		return a*k + b
	}
}
