/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package tag

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
}

func Execute() {
	fmt.Print("Add some stuff?")
}

func init() {
	TagCmd.AddCommand(AddCmd)
	TagCmd.AddCommand(ListCmd)
	TagCmd.AddCommand(ShowCmd)
}
