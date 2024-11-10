package bookmark

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("'add' expects only one argument")
		}

		request := model.BookmarkCreationRequest{Url: args[0], Title: ""}

		bookmark, dbErr := model.AddBookmark(request)
		if dbErr != nil {
			return dbErr
		}
		fmt.Println(bookmark)

		return nil
	},
}
