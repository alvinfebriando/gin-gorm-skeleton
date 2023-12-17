package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
)

func getBody(rec *httptest.ResponseRecorder) string {
	return strings.Trim(rec.Body.String(), "\n")
}

func marshal(data any) string {
	j, _ := json.Marshal(data)
	return string(j)
}

func sendBody(data any) *bytes.Reader {
	j, _ := json.Marshal(data)
	return bytes.NewReader(j)
}
