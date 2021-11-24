package cmd

import (
	"github.com/dmikalova/brocket/internal/wmctrl"

	"github.com/spf13/cobra"
)

// wmctrlCmd represents the hello command
var wmctrlCmd = &cobra.Command{
	Use:   "wmctrl",
	Short: "control windows",
	Long:  `runs the wmctrl fn`,
	Run: func(cmd *cobra.Command, args []string) {
		wmctrl.Wmctrl()
	},
}

func init() {
	rootCmd.AddCommand(wmctrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wmctrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wmctrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
