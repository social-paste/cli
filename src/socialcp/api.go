package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)

func httpRequest(method string, endpoint string, jsonBody string) (int, string) {
	url := GetHost() + endpoint

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

func UnregisterUser(email string) (int, string) {
	body := fmt.Sprintf(`{"email":"%s"}`, email)
	return httpRequest("DELETE", "accounts/delete/", body)
}

func Send(origin string, destination string, content string) (int, string) {
	body := fmt.Sprintf(`{"origin":"%s", "destination":"%s", "content":"%s"}`, origin, destination, content)
	fmt.Println(body)
	return httpRequest("POST", "messages/", body)	
}

// TODO: validar request
func Recv(origin string, destination string) (int, string) {
	body := fmt.Sprintf(`{"origin":"%s", "destination":"%s"}`, origin, destination)
	fmt.Println(body)
	return httpRequest("GET", "messages/", body)	
}
