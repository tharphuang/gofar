package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	VERSION = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show current binary version info",
	Long:    "Show current binary version info...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, "gofar CLI Tool: "+VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
