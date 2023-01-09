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
	Short: "Talks to with the API to manage Quiz",
	Long:  `Talks to with the API to manage Quiz`,
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
