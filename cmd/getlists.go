package cmd

import (
	"errors"
	"strings"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listsCmd = &cobra.Command{
	Use:   "lists FOLDERID [-a]",
	Short: "get data for lists in a folder",
	Long:  `Request JSON data for all lists by folder ID`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("incorrect number of arguments")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		checkToken()
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("archived", listCmd.Flags().Lookup("archived"))
		l := internal.ListRequest{
			FolderID: strings.Trim(args[0], " "),
			Archived: viper.GetBool("archived"),
		}
		internal.Request(l)
	},
}

func init() {
	getCmd.AddCommand(listsCmd)
	listsCmd.Flags().BoolP("archived", "a", false, "include archived lists in output")
}