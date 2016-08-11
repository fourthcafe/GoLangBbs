package db

import (
	"fmt"
	"net/url"
)

func List(url *url.Values) {
	url.Set("age", "26")

	fmt.Printf("%+v\n", url)
}
