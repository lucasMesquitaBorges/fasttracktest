/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A brief description of your command",
	Long:  `asdasdasd`,
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
