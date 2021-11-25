package cmd

import (
	"github.com/dmikalova/brocket/internal/wm"
	"github.com/spf13/cobra"
)

// runOrRaiseCmd represents the runOrRaise command
var runOrRaiseCmd = &cobra.Command{
	Use:   "run-or-raise",
	Short: "Run or raise the specified window",
	Long: `The following filters decide which window to raise:

	class: the class of the window`,
	Run: func(cmd *cobra.Command, args []string) {
		class, err := cmd.Flags().GetString("class")
		if err != nil {
			panic(err.Error())
		}
		runCmd, err := cmd.Flags().GetString("cmd")
		if err != nil {
			panic(err.Error())
		}
		list, err := cmd.Flags().GetBool("list")
		if err != nil {
			panic(err.Error())
		}

		wm.RunOrRaise(wm.Conf{
			Cmd:   runCmd,
			Class: class,
			List:  list,
		})
	},
}

func init() {
	rootCmd.AddCommand(runOrRaiseCmd)
	runOrRaiseCmd.Flags().StringP("class", "c", "", "Window class to run-or-raise")
	runOrRaiseCmd.Flags().StringP("cmd", "m", "", "Run command, defaults to class")
	runOrRaiseCmd.Flags().BoolP("list", "l", false, "List open windows and their relevant properties. The list is filtered by the other parameters.")
}
