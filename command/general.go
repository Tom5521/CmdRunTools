package command

// Execute Run on all library command types/interfaces.
func Run(cmd BaseCmd) error {
	return cmd.Run()
}

// Execute Out on all library command types/interfaces.
func Out(cmd BaseCmd) (string, error) {
	return cmd.Out()
}

// Execute CombinedOut on all library command types/interfaces.
func CombinedOut(cmd BaseCmd) (string, error) {
	return cmd.CombinedOut()
}

// Execute Start on all library command types/interfaces.
func Start(cmd BaseCmd) error {
	return cmd.Start()
}
