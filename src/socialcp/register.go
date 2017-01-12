package main

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

	return 0
}
