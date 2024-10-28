// this file is responsible for converting to json and sending json response
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Error  bool        `json:"Error"`
	Result interface{} `json:"Result"`
}

func ResultResponse(w http.ResponseWriter, statuscode int, result any, err error) {
	v := Result{
		Result: result,
	}

	encodedJson, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(statuscode)
	w.Write(encodedJson)
}
