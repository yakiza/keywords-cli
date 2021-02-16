package cli

type cliService interface {
	RunCommand(options []string) error
}

