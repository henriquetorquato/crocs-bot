package utils

import (
	"net/http"
)

var errors = make([]string, 0)
var headers = make([]http.Header, 0)

func HandleError(err error) {
	if err != nil {
		errors = append(errors, err.Error())
	}
}

func HandleResponse(response *http.Response) {
	headers = append(headers, response.Header)
}

func GetErrors() []string {
	return errors
}

func GetErrorHeaders() []http.Header {
	return headers
}
