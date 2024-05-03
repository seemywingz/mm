/*
Copyright © 2022 Kevin Jayne <kevin.jayne@icloud.com>
*/
package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	xMin,
	xMax,
	yMin,
	yMax,
	loop int

	clickClickPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mm",
	Short: "CLI Wrapper and orchestrator for CLIClick",
	Long:  `A CLI wrapper for cliclick to automate mouse movements and clicks`,
	Run: func(cmd *cobra.Command, args []string) {
		getScreenSize()
		randRClick()
		if loop > 0 {
			for {
				sleepSeconds(loop)
				randRClick()
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&loop, "loop", 0, "Toggle Automation Looping and set the delay in seconds between each iteration")
	rootCmd.Flags().IntVar(&xMin, "xmin", 0, "Minimum Horizontal Screen Address, defaults to 0")
	rootCmd.Flags().IntVar(&xMax, "xmax", -1, "Maximum Horizontal Screen Address, defaults current screen width")
	rootCmd.Flags().IntVar(&yMin, "ymin", 0, "Minimum Vertical Screen Address, defaults to 0")
	rootCmd.Flags().IntVar(&yMax, "ymax", -1, "Maximum Vertical Screen Address, defaults current screen height")
}

func sleepSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func randRClick() {
	robotgo.Move(randCoords())
	robotgo.Click("right", true)
	sleepSeconds(1)
	// get current mouse position
	x, y := robotgo.Location()
	robotgo.Move(x-10, y)
	robotgo.Click()
}

func randCoords() (int, int) {
	x := rand.Intn(xMax-xMin) + xMin
	y := rand.Intn(yMax-yMin) + yMin
	return x, y
}

func getScreenSize() {
	w, h := robotgo.GetScreenSize()
	if xMax == -1 {
		xMax = w
	}
	if yMax == -1 {
		yMax = h
	}
}
