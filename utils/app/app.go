package app

import (
	"encoding/json"
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


func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
