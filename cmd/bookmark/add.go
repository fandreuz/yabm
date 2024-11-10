package bookmark

import (
	"fmt"
	"time"

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

		url := args[0]
		creationDate := time.Now().UTC().UTC().UnixMilli()
		bookmark := model.NewBookmark(url, uint64(creationDate))

		bookmarkWithId, dbErr := model.AddBookmark(bookmark)
		if dbErr != nil {
			return dbErr
		}

		if bookmarkWithId.Id == nil {
			panic(fmt.Errorf("Unexpected nil in bookmark ID"))
		}
		fmt.Printf("Created bookmark with ID %d\n", *(bookmarkWithId.Id))
		return nil
	},
}

func init() {
}
