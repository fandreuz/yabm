package bookmark

import (
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
	Use:   "list { tagLabel | tagId } ...",
	Short: "List saved bookmarks",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		keys := tagsTrie.Search(toComplete, 5)
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		bookmarks, err := model.ListBookmarks(removeDuplicate(args))
		if err != nil {
			return err
		}

		for _, b := range bookmarks {
			fmt.Println(b.String())
		}

		return nil
	},
}
