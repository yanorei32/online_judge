package submissions

import (
	"database/sql"

	aws_session "github.com/aws/aws-sdk-go/aws/session"
)

type Submissions struct {
	DB  *sql.DB
	AWS *aws_session.Session
}
