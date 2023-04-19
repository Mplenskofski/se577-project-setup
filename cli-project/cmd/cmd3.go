/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd3Cmd represents the cmd3 command
var cmd3Cmd = &cobra.Command{
	Use:   "cmd3",
	Short: "CMD3 does something interesting",
	Long:  `CMD 3 goes out and executes some interesting web services`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd3 called")
	},
}

func init() {
	rootCmd.AddCommand(cmd3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmd3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmd3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd3Cmd.Flags().BoolP("lcmd", "l", false, "This does a boolean thing - notice I can reuse")
	cmd3Cmd.Flags().StringP("ocmd", "o", "my-o-command3", "This does a string thing - notice i can reuse")
}
