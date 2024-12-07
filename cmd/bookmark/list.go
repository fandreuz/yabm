package bookmark

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bookmarks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagNames, err := cmd.Flags().GetStringArray("tag")
		if err != nil {
			return err
		}

		bookmarks, err := model.ListBookmarks(tagNames)
		if err != nil {
			return err
		}

		for _, t := range bookmarks {
			fmt.Printf("%v\n", t)
		}

		return nil
	},
}
