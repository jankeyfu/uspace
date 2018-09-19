package controllers

import (
	"encoding/json"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty" `
}

func successRes(data interface{}) string {
	resp, err := json.Marshal(&Response{Code: "200", Msg: "", Data: data})
	if err != nil {

	}
	return string(resp)
}

func fialedRes(code, msg string) string {
	resp, err := json.Marshal(&Response{Code: code, Msg: msg})
	if err != nil {

	}
	return string(resp)
}
