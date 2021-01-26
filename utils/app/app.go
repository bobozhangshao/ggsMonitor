package app

import (
	"encoding/json"
	"ggsMonitor/utils/log"
)

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

func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return err
	}
	return nil
}
