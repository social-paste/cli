package main

import (
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

	// TODO: ler o e-mail de origem (id do usuario)
	origin := "daniela@qwerty.com"

	// TODO: traduzir o slot# para e-mail de destino
	destination := "alguem@algumlugar.com"

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

	return -1
}