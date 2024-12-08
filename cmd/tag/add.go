package tag

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/fandreuz/yabm/model/entity"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("'add' expects only one argument")
		}

		request := entity.TagCreationRequest{Label: args[0]}

		tag, dbErr := model.GetOrCreateTag(request)
		if dbErr != nil {
			return dbErr
		}
		fmt.Println(tag)

		return nil
	},
}
