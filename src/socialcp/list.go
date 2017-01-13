package main

import (
	"fmt"
)

var cmdList = &Command{
	Run:       runList,
	UsageLine: "list ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdList.Flag.BoolVar(&flagA, "a", false, "")
}

// runList executes list command and return exit code.
func runList(args []string) int {
	recipients := GetAllRecipients()

	if len(recipients) == 0 {
		fmt.Println("Nenhum destinatário cadastrado.")
		return 0
	}

	fmt.Println("|------|----------------------------------------------------|")
	fmt.Println("| Slot |                       Email                        |")
	fmt.Println("|------|----------------------------------------------------|")
	for slot, user := range recipients {
		fmt.Printf("| %4d | %50s |\n", slot, user)
	}
	fmt.Println("|------|----------------------------------------------------|")

	return 0
}
