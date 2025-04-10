package apperrors

import (
	"errors"
	"fmt"
)

var (
	ErrBeginTransaction   = errors.New("failed to begin transaction")
	ErrInsertTelegramUser = errors.New("failed to insert to telegram_user table")
	ErrInsertUser         = errors.New("failed to insert to user table")
	ErrInsertCoordinates  = errors.New("failed to insert to coordinates table")
	ErrCommitTransaction  = errors.New("failed to commit transaction")
	ErrBuildSQL           = errors.New("failed to build SQL")
	ErrExecuteQuery       = errors.New("failed to execute query")
	ErrRowsScan           = errors.New("failed to scan rows")
	ErrRows               = errors.New("failed to rows")
	ErrParseTime          = errors.New("failed to parse time")
	ErrUpdateRadius       = errors.New("failed to update radius")
	ErrUserRegistration   = errors.New("failed to registration user")
	ErrChurchGetAll       = errors.New("failed to get all churches")
)

func Wrap(err, wrapper error) error {
	return fmt.Errorf("%w: %v", wrapper, err)
}
