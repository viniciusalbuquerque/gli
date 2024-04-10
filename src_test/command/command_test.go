package command

import (
    "errors"
    "testing"

    command "github.com/viniciusalbuquerque/gli/src/command"
)

func functionToError(params []string) error {
    return errors.New("Function to error.")
}

func functionToSucceed(params []string) error {
    return nil
}

func TestCommand_ExecToFail(t *testing.T) {
    var x []string
    cmd := command.New("name", "desc", functionToError)
    err := cmd.Exec(x)

    if err == nil {
	t.Errorf("Expected function to fail")
    }
}

func TestCommand_ExecToSucceed(t *testing.T) {
    var x []string
    cmd := command.New("name", "desc", functionToSucceed)
    err := cmd.Exec(x)

    if err != nil {
	t.Errorf("Expected function to fail")
    }
}
