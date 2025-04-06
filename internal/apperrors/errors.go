package apperrors

import (
	"errors"
	"fmt"
)

var (
	ErrBeginTransaction   = errors.New("failed to begin transaction")
	ErrInsertTelegramUser = errors.New("failed to insert telegram user")
	ErrInsertUser         = errors.New("failed to insert user")
	ErrCommitTransaction  = errors.New("failed to commit transaction")
	ErrBuildSQL           = errors.New("failed to build SQL")
	ErrExecuteQuery       = errors.New("failed to execute query")
	ErrParseTime          = errors.New("failed to parse time")
)

func Wrap(err, wrapper error) error {
	return fmt.Errorf("%w: %v", wrapper, err)
}
