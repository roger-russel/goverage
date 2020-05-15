package cmd

import (
	"os"

	"github.com/roger-russel/goverage/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//CoverageFile should have where the coverage file is
var coverageFile string

//Execute the commands
func Execute(vf FullVersion) {

	checkDefaultCommand()

	rootCmd = &cobra.Command{
		Use:   "goverage [sub]",
		Short: "goverage command",
		Run: func(cmd *cobra.Command, args []string) {
			generator.BeautifulReport(cmd, args, coverageFile)
		},
	}

	rootCmd.AddCommand(Version(vf))

	rootCmd.Flags().StringVarP(&coverageFile, "covererage-file", "c", "", "Where goverage will look for the coverate report file eg: -c ./tmp/coverage.txt")
	rootCmd.MarkFlagRequired("covererage-file")

	rootCmd.Execute()

}

func checkDefaultCommand() {

	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}

}
