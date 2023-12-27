package command

func Run(cmd BaseCmd) error {
	return cmd.Run()
}

func Out(cmd BaseCmd) (string, error) {
	return cmd.Out()
}

func CombinedOut(cmd BaseCmd) (string, error) {
	return cmd.CombinedOut()
}

func Start(cmd BaseCmd) error {
	return cmd.Start()
}
