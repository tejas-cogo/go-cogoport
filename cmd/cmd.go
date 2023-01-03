package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "config",
	Short: "Application Description",
}

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
