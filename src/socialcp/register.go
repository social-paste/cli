package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
)
var cmdRegister = &Command{
	Run:       runRegister,
	UsageLine: "register ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdRegister.Flag.BoolVar(&flagA, "a", false, "")
}

// runRegister executes register command and return exit code.
func runRegister(args []string) int {
	url := "http://socialcp.cloud.dev.globoi.com/api/accounts/"

	json := []byte(fmt.Sprintf(`{"email":"%s"}`, args[0]))

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
    panic(error)
	}

	defer response.Body.Close()

	if response.StatusCode == 201 {
		fmt.Println("Usuário registrado com sucesso")
		return 0
	}

	fmt.Println("Erro ao registrar usuário")
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) != "" {
		fmt.Println("Motivo: ", string(body))
	}

	return -1
}
