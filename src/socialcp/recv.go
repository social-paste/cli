package main

import (
	"fmt"
	"strconv"
	b64  "encoding/base64"
	clip "github.com/atotto/clipboard"
	"encoding/json"
	"strings"
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

	slot, _     := strconv.Atoi(args[0])
	origin      := GetRecipient(slot)
	destination := GetUser()

	if(destination == "") {
		fmt.Println("Slot vazio.")
		return -1
	}

	// Lê o paste no servidor
	statusCode, body := Recv(origin, destination)

	// TODO: validar tratamento de erros
	if statusCode != 200 {
		fmt.Println("Erro recebendo paste.")
			if string(body) != "" {
			fmt.Println("Motivo: ", string(body))
		}
		return -1
	}

	fmt.Println("Paste recebido com sucesso.")

	fmt.Println(body)
	
	type Message struct {
	    Origin, Destination, Content string
	}

	dec := json.NewDecoder(strings.NewReader(body))
    var m Message
    if err := dec.Decode(&m); err != nil {
        panic(err)
    }

    fmt.Printf("origin: %s\ndestination: %s\n", m.Origin, m.Destination)

    msg64 := m.Content

	// Decodifica mensagem de base64 para byte[]
	message, _ := b64.StdEncoding.DecodeString(msg64)
	fmt.Println(string(message))

	// Copia mensagem para o clipboard
	if err := clip.WriteAll(string(message)); err != nil {
		panic(err)
	}

	return -1
}
