package generator

import (
	"github.com/roger-russel/goverage/internal/reader"
	"github.com/spf13/cobra"
)

//BeautifulReport takes the atomic cover profile and make a beautiful html report
func BeautifulReport(c *cobra.Command, args []string, coverageFile string) {
	data := reader.ReadFile(coverageFile)

}
