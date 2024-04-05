package command

import (
    "errors"
    "testing"

    command "github.com/viniciusalbuquerque/gli/src"
)

func functionToError(params ...any) error {
    return errors.New("Function to error.")
}

func functionToSucceed(params ...any) error {
    return nil
}

func TestCommand_ExecToFail(t *testing.T) {
    cmd := command.New("name", "desc", functionToError)
    err := cmd.Exec()

    if err == nil {
	t.Errorf("Expected function to fail")
    }
}

func TestCommand_ExecToSucceed(t *testing.T) {
    cmd := command.New("name", "desc", functionToSucceed)
    err := cmd.Exec()

    if err != nil {
	t.Errorf("Expected function to fail")
    }
}
