/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd1Cmd represents the cmd1 command
var cmd1Cmd = &cobra.Command{
	Use:   "cmd1",
	Short: "CMD1 does something interesting",
	Long:  `CMD 1 goes out and executes some interesting web services`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd1 called")
	},
}

func init() {
	rootCmd.AddCommand(cmd1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmd1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	cmd1Cmd.Flags().BoolP("lcmd", "l", false, "This does a boolean thing")
	cmd1Cmd.Flags().StringP("mcmd", "m", "my-m-command", "This does a string thing")
}
