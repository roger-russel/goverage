package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//FullVersion information
type FullVersion struct {
	Version string
	Commit  string
	Date    string
}

//Version of the binary built
func Version(vf FullVersion) (versionCmd *cobra.Command) {

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of goverage",
		Long:  `Print the semantical version of goverage built`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("v%s\nbuilded at: %s\ncommit hash: %s\n", vf.Version, vf.Date, vf.Commit)
		},
	}

	return versionCmd
}
