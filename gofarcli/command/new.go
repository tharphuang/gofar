package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create a new model of your project",
	Example: `  gofar new model`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ExecError(cmd, args, errors.New("unknown model name"))
		}
		fmt.Fprintln(os.Stdout, "gofar new: ")
	},
}

func init() {
	//newCmd.Flags().StringVarP(, "path", "p", "./", "select path")
	rootCmd.AddCommand(newCmd)
}
