package main

var cmdAdd = &Command{
	Run:       runAdd,
	UsageLine: "add ",
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

	return 0
}
