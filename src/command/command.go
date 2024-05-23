package command

import (
    "fmt"
)

type Command struct {
    name        string
    description string
    fn          func([]string) error
}

func New(name string, desc string, fn func([]string) error) *Command {
    return &Command{
        name: name,
        description: desc,
        fn: fn,
    }
}

func (cmd Command) Exec(params []string) error {
    err := cmd.fn(params)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}

func (cmd Command) PrintHelp() {
    fmt.Println("Print Help for", cmd.name)
}

