/*
Copyright © 2024 Francesco Andreuzzi <andreuzzi.francesco@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/fandreuz/yabm/cmd/bookmark"
	"github.com/fandreuz/yabm/cmd/tag"
	"github.com/fandreuz/yabm/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "yabm",
	Short:        "Simple command-line bookmark manager",
	SilenceUsage: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yabm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(bookmark.BookmarkCmd)
	bookmark.BookmarkCmd.AddCommand(MakeShowCommand(model.GetBookmarkById))
	bookmark.BookmarkCmd.AddCommand(MakeListCommand(model.ListBookmarks))

	rootCmd.AddCommand(tag.TagCmd)
	tag.TagCmd.AddCommand(MakeShowCommand(model.GetTagById))
	tag.TagCmd.AddCommand(MakeListCommand(model.ListTags))
}
