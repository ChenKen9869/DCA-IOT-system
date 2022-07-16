package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetResponseBody(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal([]byte(string(body)), &result)
	return result
}