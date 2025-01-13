package daemon

import (
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
)

func validateDaemonArgs(cmd *cobra.Command, args []string) {
	var err error
	err = config.ValidateCalculationParams()
	err = config.ValidateKey("adhan")
	if err != nil {
		util.ErrExit("%v", err)
	}
}
