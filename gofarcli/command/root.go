package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofar",
	Short: "gofar is a tool for build",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("do you forget something?")
		}
		//ExecError(cmd, args, errors.New("do you forget something"))
	},
}

func Execute() {
	rootCmd.Execute()
}
