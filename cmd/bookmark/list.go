package bookmark

import (
	"encoding/json"
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

		for _, b := range bookmarks {
			b, err := json.MarshalIndent(b, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		}

		return nil
	},
}
