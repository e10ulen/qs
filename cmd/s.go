package cmd

import (
	"os"

	"github.com/e10ulen/sandbox/lib"
	"github.com/spf13/cobra"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "markdown preview",
	Long: `markdown preview server:
	argslist
	1: port number`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		port := "PortNo :\n"
		get := lib.ScanLine(port)
		lib.MiniServe(get, dir)
	},
}

func init() {
	rootCmd.AddCommand(sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
