package users_test

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

func TestLoginSuccess(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT hashed_password FROM users WHERE name=?`)).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"hashed_password"}).AddRow("$2a$10$mfgTfkiVqozg7EItYLqp8.jGQ3KVNd9lCQNaITT5zbpEbAvXm7/su"))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, role FROM users WHERE name=?`)).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "role"}).AddRow(1, "admin", "member"))

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("{\"name\": \"admin\",\"password\": \"password\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"success\":true,\"error\":\"\",\"user\":{\"id\":1,\"name\":\"admin\",\"role\":\"member\"}}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestLoginNoUser(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT hashed_password FROM users WHERE name=?`)).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"hashed_password"}))

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("{\"name\": \"admin\",\"password\": \"password\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"specified user does not exist\",\"user\":{\"id\":0,\"name\":\"\",\"role\":\"\"}}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT hashed_password FROM users WHERE name=?`)).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"hashed_password"}).AddRow("$2a$10$mfgTfkiVqozg7EItYLqp8.jGQ3KVNd9lCQNaITT5zbpEbAvXm7/su"))

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("{\"name\": \"admin\",\"password\": \"\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"password is incorrect\",\"user\":{\"id\":0,\"name\":\"\",\"role\":\"\"}}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}
