package bookmark

import (
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete bookmarkId",
	Short: "Delete a bookmark",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return err
		}
		return model.DeleteBookmarkById(id)
	},
}
