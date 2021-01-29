package main

import (
	"fmt"
	"ggsMonitor/utils/app"
	"github.com/81120/gode/core"
	"github.com/joho/godotenv"
)

func main() {
	app.Error("加载配置文件", godotenv.Load())
	gode := core.New()
	gode.RegisterBuildInModule()

	r := gode.GetRts()
	v, err := r.RunString(`
		var r = require('./js/request.js');
		r.post(111);
	`)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(v)
}


