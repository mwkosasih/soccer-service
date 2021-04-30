package util

import (
	"google.golang.org/grpc/status"
)

// ClientError is for something that not tound
func ClientError(message string) error {
	return status.Errorf(ProcessingError, message)
}

func NotFoundError(model string) error {
	return status.Errorf(NotFound, "The "+model+" resounce not found")
}

// ErrUpdate indicates that the update is failed and need to be fixed
func ErrUpdate() error {
	return status.Errorf(InvalidArgument, "Update failed")
}

// ErrCreate indicates that the create is failed and need to be fixed
func ErrCreate() error {
	return status.Errorf(InvalidArgument, "Create failed")
}

// ErrDelete indicates that the create is failed and need to be fixed
func ErrDelete() error {
	return status.Errorf(InvalidArgument, "Delete failed")
}

func ErrProcess(message string) error {
	return status.Errorf(InvalidArgument, message)
}
