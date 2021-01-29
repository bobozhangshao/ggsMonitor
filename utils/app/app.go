package app

import (
	"ggsMonitor/utils/log"
)

type H map[string]interface{}

func Panic(err error) {
	if err != nil {
		log.PanicF("%v:", err.Error())
		panic(err.Error())
	}
}

func Error(msg string, err error) {
	if err != nil {
		log.PanicF("%s:%v:", msg, err)
	}
	log.Println("Success:", msg)
}
