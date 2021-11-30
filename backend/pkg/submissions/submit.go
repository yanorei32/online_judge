package submissions

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	echo_session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
)

type SubmitData struct {
	ProblemId int64  `json:"problem_id"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

type SubmitResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Id      int64  `json:"id"`
}

func (u *Submissions) Submit(c echo.Context) error {
	data := SubmitData{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SubmitResponse{Success: false, Error: "unknown error"})
	}
	session := echo_session.Default(c)
	login := session.Get("id")
	if login == nil || login == false {
		return c.JSON(http.StatusBadRequest, SubmitResponse{Success: false, Error: "Session is not login"})
	}
	userId := session.Get("id").(int64)
	id, err := resisterSubmitData(u.DB, userId, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SubmitResponse{Success: false, Error: err.Error()})
	}

	err = uploadCode(u.AWS, id, data.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SubmitResponse{Success: false, Error: "Failed to save the code"})
	}

	return c.JSON(http.StatusOK, SubmitResponse{Success: true, Error: "", Id: id})
}

func resisterSubmitData(db *sql.DB, userId int64, data SubmitData) (int64, error) {
	r, err := db.Exec("INSERT INTO submissions (problem_id, user_id, language) VALUES (?, ?, ?)", data.ProblemId, userId, data.Language)
	if err != nil {
		return 0, errors.New("failed to register to DB")
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, errors.New("failed to get submission id")
	}
	return id, nil
}

func uploadCode(awsSession *aws_session.Session, id int64, code string) error {
	up := aws_s3manager.NewUploader(awsSession)
	_, err := up.Upload(&aws_s3manager.UploadInput{
		Bucket:      aws.String("minio-bucket"),
		Body:        aws.ReadSeekCloser(strings.NewReader(code)),
		Key:         aws.String("submissions/" + strconv.FormatInt(id, 10)),
		ContentType: aws.String("text/plain"),
	})
	return err
}
