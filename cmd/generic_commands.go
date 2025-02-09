package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func MakeShowCommand[E fmt.Stringer](extractor func(uint64) (E, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "show id",
		Args:  cobra.ExactArgs(1),
		Short: "Show details for a specific entity",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, idConvErr := strconv.ParseUint(args[0], 10, 64)
			if idConvErr != nil {
				return idConvErr
			}

			e, dbErr := extractor(id)
			if dbErr != nil {
				return dbErr
			}

			b, jsonErr := json.Marshal(e)
			if jsonErr != nil {
				return jsonErr
			}
			fmt.Printf("%s", string(b))
			return nil
		},
	}
}
