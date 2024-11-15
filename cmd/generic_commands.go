package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func MakeShowCommand[E fmt.Stringer](extractor func(uint64) (E, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show details for a specific entity",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return fmt.Errorf("'show' expects only one argument")
			}

			id, idConvErr := strconv.ParseInt(args[0], 10, 64)
			if idConvErr != nil {
				return idConvErr
			}

			e, dbErr := extractor(uint64(id))
			if dbErr != nil {
				return dbErr
			}

			fmt.Println(e)
			return nil
		},
	}

}

func MakeListCommand[E fmt.Stringer](extractor func() ([]E, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List saved entities",
		RunE: func(cmd *cobra.Command, args []string) error {
			entities, err := extractor()
			if err != nil {
				return err
			}

			for idx, t := range entities {
				fmt.Printf("%d -- %v\n", idx, t)
			}

			return nil
		},
	}
}
