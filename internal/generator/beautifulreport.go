package generator

import (
	"fmt"

	"github.com/spf13/cobra"
)

//BeautifulReport takes the atomic cover profile and make a beautiful html report
func BeautifulReport(c *cobra.Command, args []string, coverageFile string) {
	fmt.Printf("Generate: %v\n", coverageFile)
}
