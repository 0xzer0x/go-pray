package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/util"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the CLI version",
	Long:  `The version command displays the current version of go-pray`,
	Args:  cobra.NoArgs,
	Run:   execVersion,
}

func execVersion(cmd *cobra.Command, args []string) {
	version, sum := util.Version()
	fmt.Printf("version: %s (%s)\n", version, sum)
}
