package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	// VersionCmdName represents the cobra version sub-command.
	VersionCmdName = "version"

	// VersionCmdDesp represents the cobra rest sub-command short description.
	VersionCmdDesp = "Print the version and exit"
)

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   VersionCmdName,
		Short: VersionCmdDesp,
		Run: func(*cobra.Command, []string) {
			fmt.Printf("Version: %s, BuildDate: %s", Version, BuildDate)
			os.Exit(0)
		},
	}
}
