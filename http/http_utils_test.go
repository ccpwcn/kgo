package http

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestPlainTextHttpGetRequest(t *testing.T) {
	myHttpClient := NewMyHttpClient(
		WithUrl("https://www.baidu.com"),
		WithTimeout(5*time.Second),
		WithRequestContentType(RequestPlainText),
	)
	resBodyBytes, err := myHttpClient.Get()
	if err != nil {
		t.Errorf("PlainTextHttpGetRequest failed %+v", err)
	}
	if !bytes.ContainsAny(resBodyBytes, "baidu") {
		t.Errorf("PlainTextHttpGetRequest response invalied")
	}
}

func TestJsonHttpGetRequest(t *testing.T) {
	myHttpClient := NewMyHttpClient(
		WithUrl("https://api.github.com/users/ccpwcn"),
		WithTimeout(10*time.Second),
		WithRequestContentType(RequestPlainText),
	)
	resBodyBytes, err := myHttpClient.Get()
	if err != nil {
		t.Errorf("JsonHttpGetRequest failed %+v", err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(resBodyBytes, &m)
	if err != nil {
		t.Errorf("JsonHttpGetRequest json.Unmarshal failed %+v", err)
	}
	if v, ok := m["login"]; !ok {
		t.Errorf("JsonHttpGetRequest response invalied")
	} else if vv, ok := v.(string); !ok || vv != "ccpwcn" {
		t.Errorf("JsonHttpGetRequest response invalied")
	}
}
