/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package bookmark

import (
	"github.com/spf13/cobra"
)

var BookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bookmarks",
}

func init() {
	BookmarkCmd.AddCommand(AddCmd)
}
