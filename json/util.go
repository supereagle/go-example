package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func Unmarshal2JsonObj(reader io.Reader, result interface{}) (err error) {
	if err = json.NewDecoder(reader).Decode(result); err == io.EOF {
		err = nil
		return
	}
	return
}

func UnmarshalJsonStr2Obj(jsonStr string, result interface{}) (err error) {
	if err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(result); err == io.EOF {
		err = nil
		return
	}
	return
}

func Marshal2JsonStr(jsonObj interface{}) (result string, err error) {
	bs, err := json.Marshal(jsonObj)
	if err != nil {
		return
	}
	result = bytes.NewBuffer(bs).String()
	return
}
