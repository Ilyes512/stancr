package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var stancrCmd = &cobra.Command{
	Use:   "stancr",
	Short: "A template (boilerplate) manager for scaffolding your future projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Stancr!")
	},
}

// Execute the stancr (cobra) root command
func Execute() {
	if err := stancrCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
