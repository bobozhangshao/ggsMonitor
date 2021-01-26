package app

import (
	"fmt"
	"log"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s:%v", msg, err))
	}
	log.Println("Success:", msg)
}