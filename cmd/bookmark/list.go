package bookmark

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("listtt\n")
	},
}

func init() {
}
