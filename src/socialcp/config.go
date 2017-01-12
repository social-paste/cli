package main

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

	return 0
}
