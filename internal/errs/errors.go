package errs

import "errors"

var (
	ErrUnexpectedError              = errors.New("an unexpected error happened")
	ErrCouldNotConnectDatabase      = errors.New("could not connect database")
	ErrDatabaseNotConnected         = errors.New("database were not connected")
	ErrInvalidArgumentSize          = errors.New("invalid argument size")
	ErrInvalidCommand               = errors.New("invalid command")
	ErrInvalidOperation             = errors.New("invalid operation")
	ErrInvalidArgument              = errors.New("invalid argument")
	ErrMissingInstitutionArgument   = errors.New("missing 'institution' argument")
	ErrCommandNotImplemented        = errors.New("command not implemented")
	ErrDestMustBeAddressable        = errors.New("dest argument must be addressable")
	ErrInstitutionNotFound          = errors.New("institution not found")
	ErrPlanInscriptionNotFound      = errors.New("plan inscription not found")
	ErrPlanInscriptionStoreNotFound = errors.New("plan inscription store not found")
	ErrStudentNotFound              = errors.New("student not found")
)

const (
	ExitCodeUnexpectedError = iota + 1
	ExitCodeCouldNotConnectDatabase
	ExitCodeDatabaseNotConnected
	ExitCodeInvalidArgumentSize
	ExitCodeInvalidCommand
	ExitCodeInvalidOperation
	ExitCodeInvalidArgument
	ExitCodeMissingInstitutionArgument
	ExitCodeCommandNotImplemented
	ExitCodeDestMustBeAddressable
	ExitCodeInstitutionNotFound
	ExitCodePlanInscriptionNotFound
	ExitCodePlanInscriptionStoreNotFound
	ExitCodeStudentNotFound
)

var m = map[error]int{
	ErrUnexpectedError:              ExitCodeUnexpectedError,
	ErrCouldNotConnectDatabase:      ExitCodeCouldNotConnectDatabase,
	ErrDatabaseNotConnected:         ExitCodeDatabaseNotConnected,
	ErrInvalidArgumentSize:          ExitCodeInvalidArgumentSize,
	ErrInvalidCommand:               ExitCodeInvalidCommand,
	ErrInvalidOperation:             ExitCodeInvalidOperation,
	ErrInvalidArgument:              ExitCodeInvalidArgument,
	ErrMissingInstitutionArgument:   ExitCodeMissingInstitutionArgument,
	ErrCommandNotImplemented:        ExitCodeCommandNotImplemented,
	ErrDestMustBeAddressable:        ExitCodeDestMustBeAddressable,
	ErrInstitutionNotFound:          ExitCodeInstitutionNotFound,
	ErrPlanInscriptionNotFound:      ExitCodePlanInscriptionNotFound,
	ErrPlanInscriptionStoreNotFound: ExitCodePlanInscriptionStoreNotFound,
	ErrStudentNotFound:              ExitCodeStudentNotFound,
}

func ExitCode(err error) int {
	if code, ok := m[err]; ok {
		return code
	}

	return ExitCodeUnexpectedError
}
