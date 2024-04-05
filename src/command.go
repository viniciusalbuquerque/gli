package command

import (
    "fmt"
)

type ExecFunc func(...any) error

type Command struct {
    name        string
    description string
    fn          func(...any) error
}

func New(name string, desc string, fn func(...any) error) *Command {
    return &Command{
        name: name,
        description: desc,
        fn: fn,
    }
}

func (cmd Command) Exec(params ...any) error {
    err := cmd.fn(params)
    if err != nil {
        cmd.printHelp()
        return err
    }
    return nil
}

func (cmd Command) printHelp() {
    fmt.Println("Print Help!!!")
}

