/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd2Cmd represents the cmd2 command
var cmd2Cmd = &cobra.Command{
	Use:   "cmd2",
	Short: "CMD2 does something interesting",
	Long:  `CMD 2 goes out and executes some interesting web services`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd2 called")
	},
}

func init() {
	rootCmd.AddCommand(cmd2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmd2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmd2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd2Cmd.Flags().BoolP("ncmd", "n", false, "This does a boolean thing")
	cmd2Cmd.Flags().StringP("ocmd", "o", "my-m-command-2", "This does a string thing")
}
