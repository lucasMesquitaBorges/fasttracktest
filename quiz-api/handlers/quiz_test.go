package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestQuiz(t *testing.T) {
	t.Run("Get Quiz", func(t *testing.T) {
		var quiz Quiz
		req := httptest.NewRequest("GET", "/quiz", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		l := log.New(os.Stdout, "quiz-api-test", log.LstdFlags)
		qh := NewQuiz(l)
		err := qh.StartQuiz(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("%#v", res)
		err = json.Unmarshal(res.Body.Bytes(), &quiz)
		assert.Nil(t, err)
	})
}
