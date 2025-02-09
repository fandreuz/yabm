package tag

import (
	"encoding/json"
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
			b, jsonErr := json.MarshalIndent(t, "", "  ")
			if jsonErr != nil {
				return jsonErr
			}
			fmt.Printf("%s\n", string(b))
		}

		return nil
	},
}
