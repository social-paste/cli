package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdConfig = &Command{
	Run:       runConfig,
	UsageLine: "config ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdConfig.Flag.BoolVar(&flagA, "a", false, "")
}

// runConfig executes config command and return exit code.
func runConfig(args []string) int {
	editor := "vi"

	path, error := exec.LookPath(editor)
	if error != nil {
		fmt.Println("Erro ao localizar o editor de texto " + editor + ".")
		fmt.Println(error)
		return -1
	}

	cmd := exec.Command(path, FilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	error = cmd.Start()
	if error != nil {
		fmt.Println("Erro ao iniciar o editor de texto " + editor + ".")
		fmt.Println(error)
		return -1
	}

	error = cmd.Wait()
	if error != nil {
		fmt.Println(error)
		return -1
	}

	fmt.Println("Configuração salva com sucesso!")
	return 0
}
