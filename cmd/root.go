package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qcow2file",
	Short: "qcow2file is a tool to generate qcow2 image from dockerfile",
	Long:  `qcow2file is a tool to generate qcow2 image from dockerfile`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
