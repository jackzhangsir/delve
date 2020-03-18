package proc

import (
	"errors"
	"fmt"
)

var (
	// ErrNotExecutable is returned after attempting to execute a non-executable file
	// to begin a debug session.
	ErrNotExecutable = errors.New("not an executable file")

	// ErrNotRecorded is returned when an action is requested that is
	// only possible on recorded (traced) programs.
	ErrNotRecorded = errors.New("not a recording")

	// ErrNoRuntimeAllG is returned when runtime.allg could not be found.
	ErrNoRuntimeAllG = errors.New("could not find goroutine array")
)

// ErrProcessExited indicates that the process has exited and contains both
// process id and exit status.
type ErrProcessExited struct {
	Pid    int
	Status int
}

func (pe ErrProcessExited) Error() string {
	return fmt.Sprintf("Process %d has exited with status %d", pe.Pid, pe.Status)
}

// ProcessDetachedError indicates that we detached from the target process.
type ProcessDetachedError struct{}

func (pe ProcessDetachedError) Error() string {
	return "detached from the process"
}
