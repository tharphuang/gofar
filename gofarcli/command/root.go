package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"

	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show current binary version info",
		Long:    "Show current binary version info...",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "gofar CLI Tool: "+version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(InitCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gofar",
	Short: "gofar is a tool for build",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("do you forget something?")
			ExecError(cmd, args, errors.New("do you forget something"))
		}
		//ExecError(cmd, args, errors.New("do you forget something"))
	},
}

func Execute() {
	rootCmd.Execute()
}
