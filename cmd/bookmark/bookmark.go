/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package bookmark

import (
	trie "github.com/Vivino/go-autocomplete-trie"
	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var BookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bookmarks",
}

var tagsTrie *trie.Trie = trie.New().WithoutLevenshtein().WithoutNormalisation().WithoutFuzzy().CaseSensitive()

func init() {
	BookmarkCmd.AddCommand(AddCmd)
	BookmarkCmd.AddCommand(DeleteCmd)

	BookmarkCmd.AddCommand(TagCmd)
	BookmarkCmd.AddCommand(UntagCmd)

	BookmarkCmd.AddCommand(ListCmd)

	tags, err := model.ListTags()
	if err == nil {
		tagLabels := make([]string, len(tags))
		for idx, tag := range tags {
			tagLabels[idx] = tag.Label
		}
		tagsTrie.Insert(tagLabels...)
	}

}
