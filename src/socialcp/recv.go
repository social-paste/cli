package main

import (
	"fmt"
	"strconv"
	b64  "encoding/base64"
	clip "github.com/atotto/clipboard"
)

var cmdRecv = &Command{
	Run:       runRecv,
	UsageLine: "recv slot#",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdAdd.Flag.BoolVar(&flagA, "a", false, "")
}

// runRecv executes recv command and return exit code.
func runRecv(args []string) int {

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

	// Lê o paste no servidor
	statusCode, body := Recv(origin, destination)
	fmt.Println(statusCode)

	// TODO: validar tratamento de erros
	if statusCode != 200 {
		fmt.Println("Erro recebendo paste.")
			if string(body) != "" {
			fmt.Println("Motivo: ", string(body))
		}
		return -1
	}

	fmt.Println("Paste recebido com sucesso.")
	// TODO: extrair mensagem do retorno do request
	//       (decodificar o json)
	fmt.Println(body)
	msg64 := ""

	// Decodifica mensagem de base64 para byte[]
	message, _ := b64.StdEncoding.DecodeString(msg64)
	message = []byte("abcd")
	fmt.Println(message)

	// Copia mensagem para o clipboard
	if err := clip.WriteAll(string(message)); err != nil {
		panic(err)
	}
	return -1
}
