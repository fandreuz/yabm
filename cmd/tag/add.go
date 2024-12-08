package tag

import (
	"github.com/fandreuz/yabm/model"
	"github.com/fandreuz/yabm/model/entity"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add tagLabel",
	Short: "Add a new bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		for i := range len(args) {
			request := entity.TagCreationRequest{Label: args[i]}
			_, dbErr := model.GetOrCreateTag(request)
			if dbErr != nil {
				return dbErr
			}
		}
		return nil
	},
}
