package bookmark

import (
	"fmt"
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag an existing bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("'tag' expects two arguments")
		}

		bookmarkId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if bookmarkId < 0 {
			return fmt.Errorf("Bookmark ID must be positive, got %d", bookmarkId)
		}

		tagId, err := strconv.Atoi(args[1])
		if err != nil {
			request := model.TagCreationRequest{Label: args[1]}

			tag, dbErr := model.GetOrCreateTag(request)
			if dbErr != nil {
				return dbErr
			}
			tagId = int(tag.Id)
		} else {
			if tagId < 0 {
				return fmt.Errorf("Tag ID must be positive, got %d", tagId)
			}
		}

		request := model.TagAssignationRequest{TagId: uint64(tagId), BookmarkId: uint64(bookmarkId)}
		model.AssignTag(request)

		return nil
	},
}
