package cmd

import (
	"os"

	"github.com/roger-russel/goverage/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

var flags map[string]string

//Execute the commands
func Execute(vf FullVersion) {

	checkDefaultCommand()

	rootCmd = &cobra.Command{
		Use:   "goverage [sub]",
		Short: "goverage command",
		Run: func(cmd *cobra.Command, args []string) {
			generator.BeautifulReport(cmd, args, flags)
		},
	}

	rootCmd.AddCommand(Version(vf))

	//CoverageFile should have where the coverage file is
	var coverageFile string
	rootCmd.Flags().StringVarP(&coverageFile, "coverage-file", "c", "", "Where goverage will look for the coverate report file eg: -c ./tmp/coverage.txt")
	rootCmd.MarkFlagRequired("covererage-file")

	var path string
	rootCmd.Flags().StringVarP(&path, "path", "c", "", "The root path of the project wich coverage as taken eg: -p ./go/src/myproject")

	flags["coverage-file"] = coverageFile
	flags["path"] = path

	rootCmd.Execute()

}

func checkDefaultCommand() {

	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}

}

func init() {
	flags = make(map[string]string)
}
