package submissions_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ecto0310/online_judge/backend/pkg/test_helper"
	"github.com/stretchr/testify/assert"
)

func TestSubmitSuccess(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	session, err := test_helper.CreateLoginSession(e, mock)
	if err != nil {
		t.Fatalf("Failed to create session '%#v'", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO submissions (problem_id, user_id, language) VALUES (?, ?, ?)`)).
		WithArgs(0, 1, "language").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer([]byte("{\"problem_id\": 0,\"language\": \"language\",\"code\": \"code\"}")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"success\":true,\"error\":\"\",\"id\":1}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestSubmitNotLogin(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer([]byte("{\"problem_id\": 0,\"language\": \"language\",\"code\": \"code\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"Session is not login\",\"id\":0}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}
