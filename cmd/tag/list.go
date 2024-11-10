package tag

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved bookmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tags, err := model.ListTags()
		if err != nil {
			return err
		}

		for idx, t := range tags {
			fmt.Printf("%d -- %v\n", idx, t)
		}
		
		return nil
	},
}

func init() {
}
