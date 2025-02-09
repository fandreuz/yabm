package tag

import (
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		entities, err := model.ListTags()
		if err != nil {
			return err
		}

		for _, t := range entities {
			fmt.Printf("%s\n", t.String())
		}

		return nil
	},
}
