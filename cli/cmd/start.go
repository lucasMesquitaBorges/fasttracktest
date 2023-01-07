/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new quiz",
	Long:  `Starts a new quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		startNewQuiz()
	},
}

type Quiz []Question

// Question is the structure to organize the quiz questions data
type Question struct {
	Text    string `json:"text"`
	Answer  string `json:"answer"`
	Choices []Choice
}

// Choice contains the text and value of a given questions choice
type Choice struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

// Organizes the answers given by the quizzer
type QuizAnswer struct {
	Answers []string `json:"answers"`
}

func init() {
	quizCmd.AddCommand(startCmd)
}

func promptSelectAnswer(q Question) string {
	items := []string{
		q.Choices[0].Text,
		q.Choices[1].Text,
		q.Choices[2].Text,
		q.Choices[3].Text,
	}

	var result string

	prompt := promptui.Select{
		Label: q.Text,
		Items: items,
	}

	index, result, err := prompt.Run()

	if index == -1 {
		items = append(items, result)
	}

	if err != nil {
		fmt.Printf("Prompt failed %v", err)
		os.Exit(1)
	}

	return result
}

func startNewQuiz() {
	respGet, err := http.Get("http://localhost:9090/quiz")
	if err != nil {
		fmt.Printf("Failed to get the quiz questions make sure you started the api properly")
		os.Exit(1)
	}
	defer respGet.Body.Close()
	data, _ := ioutil.ReadAll(respGet.Body)

	var r Quiz
	json.Unmarshal(data, &r)

	var answers QuizAnswer

	for _, v := range r {
		answers.Answers = append(answers.Answers, promptSelectAnswer(v))
	}

	fmt.Printf("All questions answered!\nLet's check the results!\n")

	es, err := json.Marshal(&answers)
	if err != nil {
		fmt.Printf("Failed to marshal the answers!")
		os.Exit(1)
	}
	respPost, err := http.Post("http://localhost:9090/quiz", "application/json", bytes.NewReader(es))

	if err != nil {
		fmt.Printf("Failed to get the quiz results!\nError: %v", err)
		os.Exit(1)
	}
	defer respPost.Body.Close()

	body, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s", body)
}
