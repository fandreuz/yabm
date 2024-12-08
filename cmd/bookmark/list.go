package bookmark

import (
	"encoding/json"
	"fmt"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved bookmarks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagNames, err := cmd.Flags().GetStringArray("tag")
		if err != nil {
			return err
		}

		bookmarks, err := model.ListBookmarks(removeDuplicate(tagNames))
		if err != nil {
			return err
		}

		for _, b := range bookmarks {
			b, err := json.Marshal(b)
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		}

		return nil
	},
}
