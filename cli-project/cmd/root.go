/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "se577cli",
	Short: "Demo CLI for SE577 Project",
	Long: `The purpose of this application is to create a demonstration
of a cli-based app for SE577.  This assumes your application has an
interesting architecture`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.se577cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("xflag", "x", false, "This is a global bool flag available everywhere")
	rootCmd.Flags().StringP("yflag", "y", "se577", "This is a global string flag available everywhere")
	rootCmd.Flags().Int32P("zflag", "z", 577, "This is a global int flag available everywhere")
}
