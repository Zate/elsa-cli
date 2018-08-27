package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of elsa-cli",
	Long:  `All software has versions. This software is part of ElsaCP.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("elsa-cli version " + VERSION +
            " - " + "Build date: " + SOURCE_DATE)
	},
}