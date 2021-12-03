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

func TestListSuccess(t *testing.T) {
	e, mock, err := test_helper.CreateTestServer()
	if err != nil {
		t.Fatalf("Failed to create mock server '%#v'", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT submissions.id, problem_id, problems.name, user_id, users.name, status, score, execution_time, execution_memory, submissions.created_at
			FROM (select * from submissions LIMIT 50 OFFSET ?) submissions
			LEFT JOIN users ON submissions.user_id = users.id
			LEFT JOIN problems ON submissions.problem_id = problems.id`)).
		WithArgs(0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "problem_id", "name", "user_id", "name", "status", "score", "execution_time", "execution_memory", "created_at"}).
				AddRow(1, 1, nil, 1, "admin", nil, nil, nil, nil, "2000-01-01 00:00:00"))

	req := httptest.NewRequest(http.MethodGet, "/submissions", bytes.NewBuffer([]byte("{\"problem_id\": 0,\"language\": \"language\",\"code\": \"code\"}")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"success\":true,\"error\":\"\",\"submissions\":[{\"id\":1,\"problem_id\":1,\"problem_name\":\"\",\"user_id\":1,\"user_name\":\"admin\",\"status\":\"\",\"score\":0,\"execution_time\":0,\"execution_memory\":0,\"date\":\"2000-01-01 00:00:00\"}]}\n", rec.Body.String())

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Incorrect DB call '%#v'", err)
	}
}
