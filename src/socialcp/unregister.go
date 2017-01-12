package main

var cmdUnregister = &Command{
	Run:       runUnregister,
	UsageLine: "unregister ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdUnregister.Flag.BoolVar(&flagA, "a", false, "")
}

// runUnregister executes unregister command and return exit code.
func runUnregister(args []string) int {

	return 0
}
