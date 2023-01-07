package data

import (
	"encoding/json"
	"io"
	"math"
)

// Number of questions for each quiz
const NumOfQuestion = 5

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

// In-memory variable to keep records of the amount of correct answers per user
var userAnswers = make(map[int]int, 5)

// GetQuiz returns a Quiz with the questions from the system variable quiz
func GetQuiz() Quiz {
	return quiz
}

// ToJSON encodes a Quiz structure into JSON format
func (q *Quiz) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}

// VerifyAnswers returns the number of correct answers
// It also increments the variables that serves as a record for the system
func VerifyAnswers(answers []string) int {
	// Returning variable with the number of questions answered correctly
	var crtAns int
	if len(answers) == 0 {
		return -1
	}
	// Loop through all questions
	for k, v := range quiz {
		if answers[k] == v.Answer {
			crtAns++
		}
	}

	// Incrementing systems variables that maintain records of the user answers
	userAnswers[crtAns]++

	return crtAns
}

// GetComparison compares the user results with other quizzers
func GetComparison(crtAns int) int {
	var result int
	var total int
	var uSum int

	// loop through the array of correct answers
	for k, v := range userAnswers {
		// Get's the total number of people who took the quiz
		total += v

		// Adds the amount of people that got less correct answers than the current quizzer
		if k < crtAns {
			uSum += v
		}
	}

	// Handling results
	switch {
	// If there's an error while executing the system
	case total == 0:
		result = -1

	// If the quizzer got all answers correct
	case crtAns == NumOfQuestion:
		result = int(math.Round(float64(userAnswers[NumOfQuestion]) / float64(total) * 100))

	// If the quizzer didn't get all answers correct
	default:
		result = int(math.Round(float64(uSum) / float64(total) * 100))
	}

	return result
}

// In-memory data structure for a quiz
var quiz = Quiz{
	{
		Text:   "Whats the name of the biggest country in south america?",
		Answer: "Brazil",
		Choices: []Choice{
			{
				Text:  "Argentina",
				Value: "Argentina",
			},
			{
				Text:  "Uruguay",
				Value: "Uruguay",
			},
			{
				Text:  "Brazil",
				Value: "Brazil",
			},
			{
				Text:  "Peru",
				Value: "Peru",
			},
		},
	},
	{
		Text:   "The planet Earth is the ____ planet in the solar system.",
		Answer: "Third",
		Choices: []Choice{
			{
				Text:  "First",
				Value: "First",
			},
			{
				Text:  "Second",
				Value: "Second",
			},
			{
				Text:  "Third",
				Value: "Third",
			},
			{
				Text:  "Fourth",
				Value: "Fourth",
			},
		},
	},
	{
		Text:   "What's the most venomous snake in the world?",
		Answer: "Inland Taipan",
		Choices: []Choice{
			{
				Text:  "Inland Taipan",
				Value: "Inland Taipan",
			},
			{
				Text:  "Tiger Snake",
				Value: "Tiger Snake",
			},
			{
				Text:  "King Cobra",
				Value: "King Cobra",
			},
			{
				Text:  "Saw-Scaled Viper",
				Value: "Saw-Scaled Viper",
			},
		},
	},
	{
		Text:   "What's the deadliest animal in the world?",
		Answer: "Mosquito",
		Choices: []Choice{
			{
				Text:  "Mosquito",
				Value: "Mosquito",
			},
			{
				Text:  "Snake",
				Value: "Snake",
			},
			{
				Text:  "Shark",
				Value: "Shark",
			},
			{
				Text:  "Crocodile",
				Value: "Crocodile",
			},
		},
	},
	{
		Text:   "What's was the age of the oldest person that ever lived?(years)",
		Answer: "122",
		Choices: []Choice{
			{
				Text:  "122",
				Value: "122",
			},
			{
				Text:  "104",
				Value: "104",
			},
			{
				Text:  "109",
				Value: "109",
			},
			{
				Text:  "115",
				Value: "115",
			},
		},
	},
}
