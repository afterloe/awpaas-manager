package util

import (
	"io/ioutil"
	"../../exceptions"
	"encoding/json"
	"reflect"
	"fmt"
	"net/http"
)

type responseDTO map[string]interface{}

func (res *responseDTO) String() string {
	val := reflect.ValueOf(res)
	return fmt.Sprintf("Response {data: %s, code: %d, msg: %s}", val.MapIndex(reflect.ValueOf("data")).Interface(),
		val.MapIndex(reflect.ValueOf("code")).Interface(), val.MapIndex(reflect.ValueOf("msg")).Interface())
}

func Success(data interface{}) *responseDTO {
	return Build(data, http.StatusOK, nil)
}

func Fail(code int, msg string) *responseDTO {
	return Build(nil, code, msg)
}

func Build(data interface{}, code int, msg interface{}) *responseDTO {
	return &responseDTO {
		"data": data,
		"code": code,
		"msg": msg,
	}
}

func ReadRealFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if nil != err {
		return "", &exceptions.Error{Msg: "no such this file", Code: 500}
	}
	return string(data), nil
}

func FormatToStruct(chunk *string) (map[string]interface{}, error){
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(*chunk), &rep)
	if nil != err {
		return nil, &exceptions.Error{Msg: "json format error", Code: 500}
	}
	return rep, nil
}

func FormatToString(vol interface{}) (string, error){
	buf, err := json.Marshal(vol)
	if nil != err {
		return "", &exceptions.Error{Msg: "format object error", Code: 500}
	}
	return string(buf), nil
}