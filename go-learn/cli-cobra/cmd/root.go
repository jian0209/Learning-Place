/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cli-cobra/helper"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var option bool
var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cli-cobra",
	Version: version,
	Short:   "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "This is the first command",
	Long: `A longer description
for the first command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Short:   "Reverses a string",
	Aliases: []string{"rev"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Reverse(args[0])
		fmt.Println(res)
	},
}

var uppercaseCmd = &cobra.Command{
	Use:     "uppercase",
	Short:   "Uppercase a string",
	Aliases: []string{"upper"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.UpperCase(args[0])
		fmt.Println(res)
	},
}

var modifyCmd = &cobra.Command{
	Use:     "modify",
	Short:   "Modify a string",
	Aliases: []string{"modify"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Modify(args[0], option)
		fmt.Println(res)
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	modifyCmd.Flags().BoolVarP(&option, "option", "o", false, "Modify option")
	rootCmd.AddCommand(reverseCmd)
	rootCmd.AddCommand(uppercaseCmd)
	rootCmd.AddCommand(modifyCmd)
	rootCmd.AddCommand(helloCmd)
}
