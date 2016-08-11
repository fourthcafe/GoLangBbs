package ctrl

import (
	"net/url"
	"test-duck/db"
)

func List(url *url.Values) {
	db.List(url)
}
