package bookmark

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved bookmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		bookmarks, err := model.ListBookmarks()
		if err != nil {
			return err
		}

		for _, b := range bookmarks {
			fmt.Println(b)
		}
		
		return nil
	},
}

func init() {
}
