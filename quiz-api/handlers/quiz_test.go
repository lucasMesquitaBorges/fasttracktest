package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewQuiz(t *testing.T) {
	l := log.New(os.Stdout, "", 0)
	q := NewQuiz(l)
	assert.NotNil(t, q)
}

func TestStartQuiz(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	l := log.New(os.Stdout, "", 0)
	qh := NewQuiz(l)

	// Test StartQuiz handler
	if assert.NoError(t, qh.StartQuiz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}

func TestAnswerQuiz(t *testing.T) {
	e := echo.New()
	qh := NewQuiz(log.New(os.Stdout, "", 0))

	// Test AnswerQuiz handler with empty request body
	req, err := http.NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = qh.AnswerQuiz(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "Bad Request", err.(*echo.HTTPError).Message)

	// Test AnswerQuiz handler with correct request body
	req, err = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"answers":["a","b","c"]}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = qh.AnswerQuiz(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body)

	// Test verifyRecords function with correct input
	res, err := verifyRecords(context.Background(), QuizAnswer{Answers: []string{"a", "b", "c", "d", "e"}})
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	// Test verifyRecords function with incorrect input
	_, err = verifyRecords(context.Background(), QuizAnswer{Answers: []string{}})
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	assert.Equal(t, "Couldn't fetch the user answers", err.(*echo.HTTPError).Message)
}
