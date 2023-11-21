package repository

import (
	"github.com/ansel1/merry/v2"
)

var Error = merry.Sentinel("error repository")

func MyRepository() error {

	return merry.Wrap(Error)
}
