package cmd

import (
	"clx/history"

	"github.com/spf13/cobra"
)

func clearCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clear the history of visited IDs",
		Long: "Clear the history of visited IDs from ~/.cache/circumflex/history.json.\n" +
			"History is only cached if CLX_MARK_AS_READ is set to true.",
		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			his := history.Initialize(true)
			his.ClearAndWriteToDisk()

			println("List of visited IDs cleared")
		},
	}
}
