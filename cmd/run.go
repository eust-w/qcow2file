package cmd

import (
	"fmt"
	"github.com/eust-w/qcow2file/src"
	"github.com/spf13/cobra"
	"path/filepath"
)

var qcow, currentQcow, file string
var pause bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "qcow2file is a tool to generate qcow2 image from dockerfile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		err := src.Qcow2file(qcow, currentQcow, file, pause)
		if err != nil {
			panic(err)
		}
		out, _ := filepath.Abs(currentQcow)
		fmt.Println("out qcow2 is:", out)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&qcow, "qcow", "q", "", "base qcow2 image ")
	runCmd.Flags().StringVarP(&currentQcow, "out", "o", "", "out qcow2 image ")
	runCmd.Flags().StringVarP(&file, "file", "f", "", "qcow2file like dockerfile")
	runCmd.Flags().BoolVarP(&pause, "pause", "p", false, "pause vm")
}
