package main

import (
	"github.com/gojisvm/gojis"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	AppName = "gojis"

	Profiling = "profiling"
)

var (
	rootCmd = &cobra.Command{
		Use:   AppName,
		Short: "An engine to execute ECMAScript",
		Run: func(_ *cobra.Command, args []string) {
			if viper.GetBool(Profiling) {
				p := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
				defer p.Stop()
			}

			vm := gojis.NewVM()
			_ = vm
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP(Profiling, "p", false, "Enable to write a CPU profile to \".\"")

	_ = viper.BindPFlag(Profiling, rootCmd.PersistentFlags().Lookup(Profiling))
}

func initConfig() {
	viper.AutomaticEnv()
}
