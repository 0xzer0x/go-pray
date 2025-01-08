package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/util"
	"github.com/0xzer0x/go-pray/internal/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the CLI version",
	Long:  `The version command displays the current version of go-pray`,
	Args:  cobra.NoArgs,
	Run:   execVersion,
}

func execVersion(cmd *cobra.Command, args []string) {
	semver := version.Version()
	commit := version.BuildCommit()
	buildTime, err := version.BuildTime()
	if err != nil {
		util.ErrExit("failed to extract build time: %v", err)
	}

	fmt.Printf(
		"%-15s%s\n%-15s%s\n%-15s%s\n%-15s%s\n%-15s%s\n",
		"Version:",
		semver,
		"Go Version:",
		version.Runtime(),
		"Git Commit:",
		commit,
		"Built:",
		buildTime.Format("Mon Jan 02 15:04:05 2006"),
		"OS/Arch:",
		version.OsArch(),
	)
}
