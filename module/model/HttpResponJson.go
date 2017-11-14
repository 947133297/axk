package model

import (
	"encoding/json"
	"net/http"

	"../../util"
)

type HttpResponseJson struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (thiz *HttpResponseJson) ResponseJson(w http.ResponseWriter) {
	b, err := json.Marshal(thiz)
	if err != nil {
		util.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
