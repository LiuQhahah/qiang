package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(argtemplateCmd)

}

var argtemplateCmd = &cobra.Command{
	Use:  "arg",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the arg is:" + args[0])
	},
}
