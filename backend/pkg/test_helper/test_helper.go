package test_helper

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/aws/aws-sdk-go/aws"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ecto0310/online_judge/backend/pkg/server"
	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/labstack/echo/v4"
)

func CreateTestServer() (*echo.Echo, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	redis, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}
	store, err := server.CreateSessionStore(redis.Addr(), "")
	if err != nil {
		return nil, nil, err
	}
	backend := s3mem.New()
	faker := gofakes3.New(backend)
	ts := httptest.NewServer(faker.Server())
	aws, err := server.CreateAws(ts.URL, "user", "password")
	if err != nil {
		return nil, nil, err
	}
	err = createS3Bucket(aws)
	if err != nil {
		return nil, nil, err
	}

	e := server.CreateServer(db, store)
	server.AddRouting(db, e, aws)

	return e, mock, err
}

func CreateLoginSession(e *echo.Echo, mock sqlmock.Sqlmock) (string, error) {
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

	if rec.Code != http.StatusOK {
		return "", errors.New("failed to create session")
	}
	return rec.Result().Header.Get("Set-Cookie"), nil
}

func createS3Bucket(session *aws_session.Session) error {
	client := s3.New(session)
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String("minio-bucket"),
	})
	if err != nil {
		return errors.New("failed to create bucket")
	}
	return nil
}
