package cli

import (
	"fmt"
	"go.uber.org/zap"
)

type CommandLine struct {
	cli cliService
	log zap.Logger
}

// Available options for
//	Command Line Options Available
//		--version
//		--generate { package }
// 		--instrumentation { true, false }
//

func (c CommandLine) RunCommand(options ...[]string) error {
	c.log.Info("Running commands...")
	for i := range options{
		fmt.Println(i)
	}

	return nil
}