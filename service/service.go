package service

import (
	"net/http"

	"github.com/ansel1/merry/v2"
	"github.com/edwintrumpet/merry-errors-concept/repository"
)

var Error = merry.Sentinel("error service")

func MyService() error {
	err := repository.MyRepository()

	err = merry.Apply(
		err,
		merry.WithCause(Error),
		merry.AppendMessage("es mi error"),
		merry.WithUserMessage("user not found"),
		merry.WithHTTPCode(http.StatusTeapot),
	)

	return merry.Wrap(err)
}
