/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package bookmark

import (
	"fmt"

	"github.com/spf13/cobra"
)

var BookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bookmarks",
}

func Execute() {
	fmt.Print("Add some stuff?")
}

func init() {
	BookmarkCmd.AddCommand(ListCmd)
	BookmarkCmd.AddCommand(ShowCmd)
}
