package cmd

import (
	"os"

	"github.com/roger-russel/goverage/internal/flags"
	"github.com/roger-russel/goverage/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Execute the commands
func Execute(vf FullVersion) {

	checkDefaultCommand()

	var flags flags.Flags

	rootCmd = &cobra.Command{
		Use:   "goverage [sub]",
		Short: "goverage command",
		Run: func(cmd *cobra.Command, args []string) {
			generator.BeautifulReport(cmd, args, flags)
		},
	}

	rootCmd.AddCommand(Version(vf))

	rootCmd.Flags().StringVarP(&flags.CoverageFile, "coverage-file", "c", "", "Where goverage will look for the coverate report file eg: -c ./tmp/coverage.txt")
	rootCmd.MarkFlagRequired("covererage-file")

	rootCmd.Flags().StringVarP(&flags.Path, "path", "p", "", "The root path of the project wich coverage as taken eg: -p ./go/src/myproject")
	rootCmd.Flags().StringVarP(&flags.Output, "output", "o", "", "The output file eg: -o /tmp/coverage.html")
	rootCmd.Flags().StringVar(&flags.Theme, "theme", "dracula", "The theme eg: --theme=dracula")
	rootCmd.Execute()

}

func checkDefaultCommand() {

	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}
}
