package main

import (
	"strconv"
	"fmt"
)

var cmdAdd = &Command{
	Run:       runAdd,
	UsageLine: "add ",
	Short:     "",
	Long: `

	`,
}

var cmdRemove = &Command{
	Run:       runRemove,
	UsageLine: "remove ",
	Short:     "",
	Long: `

	`,
}


func init() {
	// Set your flag here like below.
	// cmdAdd.Flag.BoolVar(&flagA, "a", false, "")
}

// runAdd executes add command and return exit code.
func runAdd(args []string) int {
	slot, _ := strconv.Atoi(args[0])
	dest := args[1]
	SetRecipient(slot, dest)
	fmt.Println("Usuário", dest, "adicionado com êxito ao slot", slot)
	return 0
}

func runRemove(args []string) int {
	slot, _ := strconv.Atoi(args[0])
	RemoveRecipient(slot)
	fmt.Println("Usuário removido do slot", slot)
	return 0
}
