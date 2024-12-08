package tag

import (
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete { tagLabel | tagId }",
	Short: "Delete a tag",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		tagId, err := strconv.ParseUint(args[0], 10, 64)
		if err == nil {
			return model.DeleteTagById(tagId)
		} else {
			return model.DeleteTagByLabel(args[0])
		}
	},
}
