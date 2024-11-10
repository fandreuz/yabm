package bookmark

import (
	"fmt"
	"strconv"

	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show bookmark details",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("'show' expects only one argument")
		}

		id, idConvErr := strconv.ParseInt(args[0], 10, 64)
		if idConvErr != nil {
			return idConvErr
		}

		bookmark, dbErr := model.GetBookmarkById(uint64(id))
		if dbErr != nil {
			return dbErr
		}

		fmt.Println(bookmark)
		return nil
	},
}
