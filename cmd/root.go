package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qiang",
	Short: "",
	Long:  "",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}
