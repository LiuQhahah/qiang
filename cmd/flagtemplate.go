package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flagCmd)
	flagCmd.Flags().StringP("flag", "f", "flag", "flag template")
}

var flagCmd = &cobra.Command{
	Use: "flag",
	Run: func(cmd *cobra.Command, args []string) {
		flagName, _ := cmd.Flags().GetString("flag")
		fmt.Println("the value of flag is :" + flagName)
	},
}
