package main

import (
	"fmt"
)

var cmdRegister = &Command{
	Run:       runRegister,
	UsageLine: "register ",
	Short:     "",
	Long: `

	`,
}

var cmdUnregister = &Command{
	Run:       runUnregister,
	UsageLine: "unregister ",
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
	statusCode, body := RegisterUser(args[0])

	if statusCode == 201 {

		SetUser(args[0])
		fmt.Println("Usuário registrado com sucesso")
		return 0
	}

	if string(body) != "" {
		fmt.Println("Motivo: ", string(body))
	}

	return -1
}

func runUnregister(args []string) int {
	statusCode, body := UnregisterUser(args[0])

	if statusCode == 204 {
		fmt.Println("Usuário removido com sucesso")
		return 0
	}

	if string(body) != "" {
		fmt.Println("Motivo: ", string(body))
	}

	return -1
}
