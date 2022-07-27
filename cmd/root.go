/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	xmin,
	xmax,
	ymin,
	ymax int
	loop bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mm",
	Short: "CLI Wrapper and orcestrator for CLIClick",
	Long:  `A CLI wrapper for cliclick to automate mouse movements and clicks`,
	Run: func(cmd *cobra.Command, args []string) {
		randRC()
		sleepSeconds(3)
		for loop {
			randRC()
			sleepSeconds(3)
		}
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
	rootCmd.Flags().BoolVar(&loop, "loop", false, "Toggle Automation Looping")
	rootCmd.Flags().IntVar(&xmin, "xmin", 9, "Minimum Horizontal Screen Address")
	rootCmd.Flags().IntVar(&xmax, "xmax", 999, "Maximum Horizontal Screen Address")
	rootCmd.Flags().IntVar(&ymin, "ymin", 9, "Minimum Vertical Screen Address")
	rootCmd.Flags().IntVar(&ymax, "ymax", 999, "Maximum Vertical Screen Address")
}
