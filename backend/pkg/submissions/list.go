package submissions

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SubmissionsData struct {
	Id              int64  `json:"id"`
	ProblemId       int64  `json:"problem_id"`
	ProblemName     string `json:"problem_name"`
	UserId          int64  `json:"user_id"`
	UserName        string `json:"user_name"`
	Status          string `json:"status"`
	Score           int    `json:"score"`
	ExecutionTime   int64  `json:"execution_time"`
	ExecutionMemory int64  `json:"execution_memory"`
	Date            string `json:"date"`
}

type SubmissionListResponse struct {
	Success     bool              `json:"success"`
	Error       string            `json:"error"`
	Submissions []SubmissionsData `json:"submissions"`
}

func (u *Submissions) List(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		page = 0
	}

	submissionsList, err := getSubmissionsList(u.DB, page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SubmissionListResponse{Success: false, Error: err.Error()})
	}

	return c.JSON(http.StatusOK, SubmissionListResponse{Success: true, Submissions: submissionsList})
}

func getSubmissionsList(db *sql.DB, page int64) ([]SubmissionsData, error) {
	rows, err := db.Query(
		`SELECT submissions.id, problem_id, problems.name, user_id, users.name, status, score, execution_time, execution_memory, submissions.created_at
			FROM (select * from submissions LIMIT 50 OFFSET ?) submissions
			LEFT JOIN users ON submissions.user_id = users.id
			LEFT JOIN problems ON submissions.problem_id = problems.id`, page*50)
	if err != nil {
		return nil, errors.New("failed to get submission list")
	}
	submissions := []SubmissionsData{}
	for rows.Next() {
		submission := struct {
			Id              int64
			ProblemId       int64
			ProblemName     sql.NullString
			UserId          int64
			UserName        sql.NullString
			Status          sql.NullString
			Score           sql.NullInt32
			ExecutionTime   sql.NullInt64
			ExecutionMemory sql.NullInt64
			Date            string
		}{}
		err := rows.Scan(&submission.Id, &submission.ProblemId, &submission.ProblemName, &submission.UserId, &submission.UserName,
			&submission.Status, &submission.Score, &submission.ExecutionTime, &submission.ExecutionMemory, &submission.Date)
		if err != nil {
			return nil, errors.New("unknown error")
		}
		submissions = append(submissions, SubmissionsData{
			Id:              submission.Id,
			ProblemId:       submission.ProblemId,
			ProblemName:     submission.ProblemName.String,
			UserId:          submission.UserId,
			UserName:        submission.UserName.String,
			Status:          submission.Status.String,
			Score:           int(submission.Score.Int32),
			ExecutionTime:   submission.ExecutionTime.Int64,
			ExecutionMemory: submission.ExecutionMemory.Int64,
			Date:            submission.Date,
		})
	}
	return submissions, nil
}
