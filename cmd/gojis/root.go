package main

import (
	"github.com/gojisvm/gojis"
	"github.com/gojisvm/gojis/config"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// AppName is the name of the executable CLI application, that will be used
	// in the help text.
	AppName = "gojis"

	// debug is the constant key for the flag and configuration entry "debug"
	debug          = "debug"
	debugShorthand = "d"
	debugDefault   = false

	// profiling is the constant key for the flag and configuration entry
	// "profiling".
	profiling          = "profiling"
	profilingShorthand = "p"
	profilingDefault   = false

	// profilePath is the constant key for the flag and configuration entry
	// "profile-path", which determines the location of the written CPU profile.
	profilePath        = "profile-path"
	profilePathDefault = "."
)

var (
	rootCmd = &cobra.Command{
		Use:   AppName,
		Short: "An engine to execute ECMAScript",
		Run: func(_ *cobra.Command, args []string) {
			if viper.GetBool(profiling) {
				p := profile.Start(profile.CPUProfile, profile.ProfilePath(viper.GetString(profilePath)))
				defer p.Stop()
			}

			vm := gojis.NewVM(&config.Cfg{
				Debug: viper.GetBool(debug),
			})
			_ = vm
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP(debug, debugShorthand, debugDefault, "Enable only if you know what you're doing")
	rootCmd.PersistentFlags().BoolP(profiling, profilingShorthand, profilingDefault, "Enable to write a CPU profile")
	rootCmd.PersistentFlags().String(profilePath, profilePathDefault, "Set the output path for the CPU profile. Only has an effect if profiling is enabled")

	_ = viper.BindPFlag(debug, rootCmd.PersistentFlags().Lookup(debug))
	_ = viper.BindPFlag(profiling, rootCmd.PersistentFlags().Lookup(profiling))
	_ = viper.BindPFlag(profilePath, rootCmd.PersistentFlags().Lookup(profilePath))
}

func initConfig() {
	viper.AutomaticEnv()
}
