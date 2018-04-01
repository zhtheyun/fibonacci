package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// RootCmdName represents the cobra root command.
	RootCmdName = "fibonacci"

	// RootCmdDesp represents the cobra root command description.
	RootCmdDesp = "An application for fibonacci"
)

// Version represents the binary version.
var Version string

// BuildDate represents the binary build date.
var BuildDate string

// init just follows the cobra convention.
func init() {
	cobra.OnInitialize(initViper)
}

// initViper init viper configuration.
func initViper() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// It follows the cobra convention.
func Execute(version, builddate string) {
	Version = version
	BuildDate = builddate
	viper.Set("VERSION", version)
	viper.Set("BUILDDATE", builddate)
	rootCmd := &cobra.Command{
		Use:   RootCmdName,
		Short: RootCmdDesp,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	rootCmd.AddCommand(commandVersion())
	rootCmd.AddCommand(commandRest())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute command: %+v\n", err)
		os.Exit(1)
	}
}
