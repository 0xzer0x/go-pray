package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
	"github.com/0xzer0x/go-pray/internal/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  `Display the current version information of go-pray`,
	Args:  cobra.NoArgs,
	Run:   execVersion,
}

func execVersion(cmd *cobra.Command, args []string) {
	versionInfo, err := version.NewVersionInfo()
	if err != nil {
		util.ErrExit("%v", err)
	}

	formatter := config.Formatter()
	output, err := formatter.VersionInfo(versionInfo)
	if err != nil {
		util.ErrExit("%v", err)
	}

	fmt.Print(output)
}
