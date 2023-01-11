package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alexralbino/fasttracktest/quiz-api/data"
	"github.com/labstack/echo/v4"
)

// Quiz is a http.Handler
type Quiz struct {
	l *log.Logger
}

// Organizes the answers given by the quizzer
type QuizAnswer struct {
	Answers []string `json:"answers"`
}

// Creates a new Quiz handler with the given logger
func NewQuiz(l *log.Logger) *Quiz {
	return &Quiz{l}
}

// StartQuiz is a http.HandleFunc that retrieves a new Quiz and returns to the user
func (q *Quiz) StartQuiz(c echo.Context) error {
	// Fetch the questions list for the quiz
	lq := data.GetQuiz()

	// Serialize the list into JSON
	err := lq.ToJSON(c.Response().Writer)

	if err != nil {
		http.Error(c.Response().Writer, "Unable to unmarshal the quiz", http.StatusInternalServerError)
	}
	return nil
}

// AnswerQuiz is a http.HandleFunc that posts the user input.
func (q *Quiz) AnswerQuiz(c echo.Context) error {
	var quizAns QuizAnswer

	// Binds the answer into a QuizAnswer struct
	err := c.Bind(&quizAns)
	if err != nil {
		return err
	}

	if quizAns.Answers == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	strResult, err := verifyRecords(context.Background(), quizAns)

	return c.String(http.StatusOK, strResult)
}

func verifyRecords(ctx context.Context, q QuizAnswer) (string, *echo.HTTPError) {
	// Checks how many answers the quizzer answered correctly
	crtAns := data.VerifyAnswers(q.Answers)
	if crtAns == -1 {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Couldn't fetch the user answers")
	}
	// Compares the quizzer with others
	comp := data.GetComparison(crtAns)

	// Checks for error while processing the comparison
	if comp == -1 {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Error while processing your answers")
	}

	// Pretty printing the result to the quizzer
	//
	//
	strResult := fmt.Sprintf("You correctly answered %d of the questions!\n", crtAns)
	if data.NumOfQuestion == crtAns {
		strResult += fmt.Sprintf("Congratulations, that's all %d questions correctly answered.\nYou're in the %d%% of quizzers that got all questions right!", data.NumOfQuestion, comp)
	} else {
		strResult += fmt.Sprintf("You were better than %d%% of all quizzers!!", comp)
	}
	return strResult, nil
}
