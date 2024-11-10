package tag

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

		label := args[0]
		creationDate := time.Now().UTC().UTC().UnixMilli()
		tag := model.NewTag(label, uint64(creationDate))

		tagWithId, dbErr := model.AddTag(tag)
		if dbErr != nil {
			return dbErr
		}
		if tagWithId.Id == nil {
			panic(fmt.Errorf("Unexpected nil in tag ID"))
		}

		fmt.Printf("Created tag with ID %d\n", *(tagWithId.Id))
		return nil
	},
}

func init() {
}
