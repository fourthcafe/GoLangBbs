package go_routine

import (
	"fmt"
	"math/rand"
	"time"
)

func Main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	fmt.Scanln()
}

func hello(i int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(i)
}
