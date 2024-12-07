/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package tag

import (
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
}

func init() {
	TagCmd.AddCommand(AddCmd)
	TagCmd.AddCommand(ListCmd)
}
