package main

import (
	"errors"
	"fmt"

	"github.com/ansel1/merry/v2"
	"github.com/edwintrumpet/merry-errors-concept/repository"
	"github.com/edwintrumpet/merry-errors-concept/service"
)

func main() {
	err := service.MyService()

	if err == nil {
		fmt.Println("no error")
		return
	}

	if err == service.Error {
		fmt.Println("error is equal to error service")
	}

	if err == repository.Error {
		fmt.Println("error is equal to error repository")
	}

	if errors.Is(err, service.Error) {
		fmt.Println("error matches with error service")
	}

	if errors.Is(err, repository.Error) {
		fmt.Println("error matches with error repository")
	}

	fmt.Println("err:", err)

	stack := merry.Stacktrace(err)
	userMessage := merry.UserMessage(err)
	httpCode := merry.HTTPCode(err)

	fmt.Println("stack:", stack)
	fmt.Println("userMessage:", userMessage)
	fmt.Println("httpCode:", httpCode)
}
