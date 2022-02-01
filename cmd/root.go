package cmd

import (
	"fmt"
	"os"

	"bloom/config"

	"github.com/spf13/cobra"
)

var (
	out string

	rootCmd = &cobra.Command{
		Use:     config.AppName,
		Short:   config.AppDesc,
		Long:    config.AppDesc,
		Version: config.AppVersion,
	}
)

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {

}
