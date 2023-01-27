package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-guardian/guardian"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome     = "home"
	FlagLoglevel = "loglevel"

	Version = "dev"
)

func init() {
	rootCmd.PersistentFlags().String(FlagHome, os.ExpandEnv("$HOME/.im-guardian"), "home path")
	rootCmd.PersistentFlags().String(FlagLoglevel, "info", "log level")

	rootCmd.AddCommand(
		CmdVersion(),
		CmdStart(),
	)
}

var rootCmd = &cobra.Command{
	Use: "im-guardian",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		home, err := cmd.Flags().GetString(FlagHome)
		CheckErr(err)
		viper.SetDefault(FlagHome, home)

		ll, err := cmd.Flags().GetString(FlagLoglevel)
		CheckErr(err)
		log.SetLevelByName(ll)

		setupConfig(home)
		return nil
	},
}

func CmdStart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start message app guardian service",
		Run: func(cmd *cobra.Command, args []string) {
			guardian.InitChains()
			guardian.Start()
		},
	}
	return cmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setupConfig(home string) {
	readConfig(home, "config/config.toml")
}

func readConfig(home, relativePath string) {
	path := filepath.Join(home, relativePath)
	viper.SetConfigFile(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("failed to load", path, err)
	}
}

func CmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Version of the binary",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
	return cmd
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("fatal error occurred: %s", err.Error())
	}
}
