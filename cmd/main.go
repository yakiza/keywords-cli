package cmd

import (
	"fmt"
	"github.com/yakiza/keywords-cli/cmd/cli"
	"os"
)

func main(){
	fmt.Println("Welcome to this tool")
	cmd := cli.CommandLine{}

	cmd.RunCommand(os.Args)
}
