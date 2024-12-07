package bookmark

import (
	"fmt"
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var UntagCmd = &cobra.Command{
	Use:   "untag",
	Short: "Untag a bookmark",
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
			request := model.TagAssignationByLabelRequest{TagLabel: args[1], BookmarkId: bookmarkId}
			return model.UnassignTagByLabel(request)
		} else {
			request := model.TagAssignationRequest{TagId: tagId, BookmarkId: bookmarkId}
			return model.UnassignTag(request)
		}
	},
}
