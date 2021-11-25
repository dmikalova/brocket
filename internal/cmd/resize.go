package cmd

import (
	"github.com/dmikalova/brocket/internal/brocket"

	"github.com/spf13/cobra"
)

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize windows",
	Long: `Use brocket to resize and place windows.

	All values are 0-100/100 multipliers to the max width and height of the desktop.`,
	Run: func(cmd *cobra.Command, args []string) {
		// f, err := cmd.Flags().GetBool("frame")
		// if err != nil {
		// 	panic(err.Error())
		// }
		h, err := cmd.Flags().GetInt("height")
		if err != nil {
			panic(err.Error())
		}
		w, err := cmd.Flags().GetInt("width")
		if err != nil {
			panic(err.Error())
		}
		x, err := cmd.Flags().GetInt("x-offset")
		if err != nil {
			panic(err.Error())
		}
		y, err := cmd.Flags().GetInt("y-offset")
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println("flags", f)
		c := brocket.Conf{
			// Frame:  f,
			Height: h,
			Width:  w,
			X:      x,
			Y:      y,
		}

		brocket.Resize(c)
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)

	// resizeCmd.Flags().BoolP("frame", "f", false, "Offset with window frame")
	resizeCmd.Flags().IntP("height", "t", 100, "Window height")
	resizeCmd.Flags().IntP("width", "w", 100, "Window width")
	resizeCmd.Flags().IntP("x-offset", "x", 0, "Window x offset")
	resizeCmd.Flags().IntP("y-offset", "y", 0, "Window y offset")

}
