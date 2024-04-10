package gli

import (
	"errors"
	"testing"

	gli "github.com/viniciusalbuquerque/gli/src/gli"
)

func fnToPass(pms []string) error {
    return nil
}

func fnToError(pms []string) error {
    return errors.New("Error")
}

func TestGli_ExecCommandNotEnoughArgs(t *testing.T) {
    gli.RegisterCommand("name", "desc", fnToError)
    
    var x []string

    err := gli.ExecCommand(x)

    if err == nil {
	t.Errorf("Expected it to error due to passing no arguments")
    }
}

func TestGli_ExecCommandCommandNotRegistered(t *testing.T) {
    gli.RegisterCommand("name", "desc", fnToError)
    
    x := []string{"anything", "unnamed", "whatever"}
    err := gli.ExecCommand(x)

    if err == nil {
	t.Errorf("It should fail due to trying to execute an unregisterd command")
    }
}

func TestGli_ExecCommandCommandToPass(t *testing.T) {
    gli.RegisterCommand("name", "desc", fnToPass)
    
    x := []string{"anything", "name", "whatever"}
    err := gli.ExecCommand(x)

    if err != nil {
	t.Errorf("It should have passed")
    }
}

func TestGli_ExecCommandCommandToFail(t *testing.T) {
    gli.RegisterCommand("name", "desc", fnToError)
    
    x := []string{"anything", "name", "whatever"}
    err := gli.ExecCommand(x)

    if err == nil {
	t.Errorf("It should have erroed")
    }
}

func TestGli_ExecCommandCommandWhenNoCommandsWereRegistered(t *testing.T) {
    x := []string{"anything", "name", "whatever"}
    err := gli.ExecCommand(x)

    if err == nil {
	t.Errorf("It should have erroed due to no commands were ever registered")
    }
}
