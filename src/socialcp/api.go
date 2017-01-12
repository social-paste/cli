package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)

func httpRequest(method string, endpoint string, jsonBody string) (int, string) {
	url := "http://socialcp.cloud.dev.globoi.com/api/" + endpoint

	json := []byte(jsonBody)

	request, error := http.NewRequest(method, url, bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
    panic(error)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return response.StatusCode, string(body)
}

func RegisterUser(email string) (int, string) {
	body := fmt.Sprintf(`{"email":"%s"}`, email)
	return httpRequest("POST", "accounts/", body)
}
