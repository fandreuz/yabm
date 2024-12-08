package bookmark

import (
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/fandreuz/yabm/model/entity"
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag bookmarkId { tagLabel | tagId } ...",
	Short: "Tag a bookmark",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bookmarkId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return err
		}

		for i := 1; i < len(args); i++ {
			var err error
			tagId, err := strconv.ParseUint(args[i], 10, 64)
			if err == nil {
				request := entity.TagAssignationRequest{TagId: tagId, BookmarkId: bookmarkId}
				_, err = model.AssignTagById(request)
			} else {
				request := entity.TagAssignationByLabelRequest{TagLabel: args[i], BookmarkId: bookmarkId}
				_, err = model.AssignTagByLabel(request)
			}

			if err != nil {
				return err
			}
		}

		return nil
	},
}
