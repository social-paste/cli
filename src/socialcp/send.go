package main

import (
	"strconv"
	"fmt"
	b64  "encoding/base64"
	clip "github.com/atotto/clipboard"
)

var cmdSend = &Command{
	Run:       runSend,
	UsageLine: "send slot#",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdAdd.Flag.BoolVar(&flagA, "a", false, "")
}

// runSend executes send command and return exit code.
func runSend(args []string) int {

	// TODO: validar o parametro
	if( len(args) == 0 ) {
		fmt.Println("Comando mal formado: especifique o slot#")
		return -1
	}

	slot, _ := strconv.Atoi(args[0])
	origin := GetUser()
	destination := GetRecipient(slot)

	if(destination == "") {
		fmt.Println("Slot vazio.")
		return -1
	}

	// Carrega a mensagem do clipboard
	message, err := clip.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(message)

	// Codifica para base64 para poder transferir caracteres especiais
	msg64 := b64.StdEncoding.EncodeToString([]byte(message))
	statusCode, body := Send(origin, destination, msg64)

	if statusCode == 201 {
		fmt.Println("Enviado com sucesso.")
		return 0
	}

	if string(body) != "" {
		fmt.Println("Motivo: ", string(body))
	}

	return 0
}
