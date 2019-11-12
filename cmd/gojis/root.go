package main

import (
	"github.com/gojisvm/gojis"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// AppName is the name of the executable CLI application, that will be used
	// in the help text.
	AppName = "gojis"

	// Profiling is the constant key for the flag and configuration entry
	// "profiling".
	Profiling          = "profiling"
	ProfilingShorthand = "p"
	ProfilingDefault   = false

	// ProfilePath is the constant key for the flag and configuration entry
	// "profile-path", which determines the location of the written CPU profile.
	ProfilePath        = "profile-path"
	ProfilePathDefault = "."
)

var (
	rootCmd = &cobra.Command{
		Use:   AppName,
		Short: "An engine to execute ECMAScript",
		Run: func(_ *cobra.Command, args []string) {
			if viper.GetBool(Profiling) {
				p := profile.Start(profile.CPUProfile, profile.ProfilePath(viper.GetString(ProfilePath)))
				defer p.Stop()
			}

			vm := gojis.NewVM()
			_ = vm
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP(Profiling, ProfilingShorthand, ProfilingDefault, "Enable to write a CPU profile")
	rootCmd.PersistentFlags().String(ProfilePath, ProfilePathDefault, "Set the output path for the CPU profile. Only has an effect if profiling is enabled")

	_ = viper.BindPFlag(Profiling, rootCmd.PersistentFlags().Lookup(Profiling))
	_ = viper.BindPFlag(ProfilePath, rootCmd.PersistentFlags().Lookup(ProfilePath))
}

func initConfig() {
	viper.AutomaticEnv()
}
