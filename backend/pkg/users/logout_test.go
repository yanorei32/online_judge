package users_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ecto0310/online_judge/backend/pkg/test_helper"
	"github.com/stretchr/testify/assert"
)

func TestLogoutSuccess(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	session, err := test_helper.CreateLoginSession(e, mock)
	if err != nil {
		t.Fatalf("Failed to create session '%#v'", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/logout", bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", session)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"success\":true,\"error\":\"\"}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestLogoutNotLogin(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/logout", bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"Session is not login\"}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}
