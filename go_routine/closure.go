package go_routine

import (
	"fmt"
	"runtime"
)

func MainClosure() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world!"

	// 고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행
	// 바뀌는 변수는 반드시 매개변수로 넘겨주기
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
