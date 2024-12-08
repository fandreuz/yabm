package bookmark

import (
	"fmt"
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/fandreuz/yabm/model/entity"
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag a bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("'tag' expects two arguments")
		}

		bookmarkId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return err
		}
		if bookmarkId < 0 {
			return fmt.Errorf("Bookmark ID must be positive, got %d", bookmarkId)
		}

		tagId, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			request := entity.TagCreationRequest{Label: args[1]}

			tag, dbErr := model.GetOrCreateTag(request)
			if dbErr != nil {
				return dbErr
			}
			tagId = tag.Id
		} else {
			if tagId < 0 {
				return fmt.Errorf("Tag ID must be positive, got %d", tagId)
			}
		}

		request := entity.TagAssignationRequest{TagId: tagId, BookmarkId: bookmarkId}
		model.AssignTag(request)

		return nil
	},
}
