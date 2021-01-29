package request

import (
	"ggsMonitor/utils/app"
)

// response http response wrapper
type response struct {
	body []byte
}

// String return response as string
func (wrap *response) String() string {
	return string(wrap.body)
}

// BindJson bind json object with pointer
func (wrap *response) BindJson(object interface{}) {
	app.JsonToStruct(string(wrap.body), object)
}