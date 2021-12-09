package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func ExecuteHelper(name string, subName string, args ...string) (string, error) {
	args = append([]string{subName}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

func ExecError(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "gofar %s %v: %s\n ", cmd.Name(), args, err.Error())
}

func CommandError(cmd *cobra.Command) {
	fmt.Println(cmd.Name())
}
