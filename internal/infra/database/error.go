package database

import "fmt"

func newPreparationErr(queryName string, repository string, err error) *DatabaseErr {
	return &DatabaseErr{
		Message: fmt.Sprintf("unable to prepare the query:`%s` on %s repository, original err: %s", queryName, repository, err.Error()),
	}
}

func newStatementNotPreparedErr(queryName string, repository string) *DatabaseErr {
	return &DatabaseErr{
		Message: fmt.Sprintf("query `%s` is not prepared on %s repository", queryName, repository),
	}
}

type DatabaseErr struct {
	Message string
}

func (e *DatabaseErr) Error() string {
	return e.Message
}
