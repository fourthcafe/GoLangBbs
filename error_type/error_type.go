package error_type

import (
	"fmt"
	"log"
	"time"
)

func Main() {
	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	s, err = helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	fmt.Println("It`s successful end...")
}

type HelloOneError struct {
	time  time.Time
	value int
}

func (this HelloOneError) Error() string {
	return fmt.Sprintf("%v: %d 는 1이 아님\n", this.time, this.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", HelloOneError{time.Now(), n}
}
