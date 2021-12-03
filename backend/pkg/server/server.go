package server

import (
	"database/sql"

	"github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/ecto0310/online_judge/backend/pkg/submissions"
	"github.com/ecto0310/online_judge/backend/pkg/users"
	_ "github.com/go-sql-driver/mysql"
	echo_session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateDbConnection(address string) (*sql.DB, error) {
	db, err := sql.Open("mysql", address)
	return db, err
}

func CreateSessionStore(address string, password string) (echo_session.RedisStore, error) {
	store, err := echo_session.NewRedisStore(32, "tcp", address, password, make([]byte, 32))
	if err != nil {
		return nil, err
	}
	store.Options(echo_session.Options{Path: "/", MaxAge: 86400 * 7})
	return store, nil
}

func CreateAws(address string, id string, secret string) (*aws_session.Session, error) {
	sess, err := aws_session.NewSessionWithOptions(aws_session.Options{
		Config: aws.Config{
			Credentials:      aws_credentials.NewStaticCredentials(id, secret, ""),
			Endpoint:         aws.String(address),
			Region:           aws.String("ap-northeast-1"),
			S3ForcePathStyle: aws.Bool(true),
		},
		Profile: "default",
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func CreateServer(db *sql.DB, store echo_session.RedisStore) *echo.Echo {
	r := echo.New()

	r.Pre(middleware.RemoveTrailingSlash())

	r.Use(echo_session.Sessions("SESSION", store))
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	return r
}

func AddRouting(db *sql.DB, s *echo.Echo, aws *aws_session.Session) {
	users := &users.Users{DB: db}
	s.POST("/register", users.Register)
	s.POST("/login", users.Login)
	s.GET("/logout", users.Logout)

	submissions := &submissions.Submissions{DB: db, AWS: aws}
	s.POST("/submit", submissions.Submit)
	s.GET("/submissions", submissions.List)
}
