package main

var cmdRemove = &Command{
	Run:       runRemove,
	UsageLine: "remove ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdRemove.Flag.BoolVar(&flagA, "a", false, "")
}

// runRemove executes remove command and return exit code.
func runRemove(args []string) int {

	return 0
}
