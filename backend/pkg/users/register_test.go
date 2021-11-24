package users_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ecto0310/online_judge/backend/pkg/test_helper"
	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (email, name, hashed_password) VALUES (?, ?, ?)`)).
		WithArgs("admin@example.com", "admin", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, role FROM users WHERE name=?`)).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "role"}).AddRow(1, "admin", "member"))

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("{\"email\": \"admin@example.com\",\"name\": \"admin\",\"password\": \"password\"}")))
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

func TestRegisterNameValidate(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("{\"email\": \"admin@example.com\",\"name\": \"\",\"password\": \"password\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"insufficient name length\",\"user\":{\"id\":0,\"name\":\"\",\"role\":\"\"}}\n", rec.Body.String())
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestRegisterPasswordValidate(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("{\"email\": \"admin@example.com\",\"name\": \"admin\",\"password\": \"\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"insufficient password length\",\"user\":{\"id\":0,\"name\":\"\",\"role\":\"\"}}\n", rec.Body.String())
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}

func TestRegisterDuplicate(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (email, name, hashed_password) VALUES (?, ?, ?)`)).
		WithArgs("admin@example.com", "admin", sqlmock.AnyArg()).
		WillReturnError(fmt.Errorf(""))

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("{\"email\": \"admin@example.com\",\"name\": \"admin\",\"password\": \"password\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"success\":false,\"error\":\"failed to register to DB\",\"user\":{\"id\":0,\"name\":\"\",\"role\":\"\"}}\n", rec.Body.String())
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}
