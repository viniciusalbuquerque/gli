package gli

import (
	"fmt"

	"github.com/viniciusalbuquerque/gli/src/command"
)

var m map[string]*command.Command

func RegisterCommand(name string, desc string, fn func([]string) error) {

    if m == nil {
        m = make(map[string]*command.Command)
    }

    m[name] = command.New(name, desc, fn)
}

func ExecCommand(args []string) error {
    if len(args) < 2  {
        printHelp()
        return fmt.Errorf("Not enough arguments.")
    }

    if m == nil || len(m) == 0 {
        return fmt.Errorf("No commands were registered")
    }

    cmd := args[1]
    if m[cmd] == nil {
        printHelp()
        return fmt.Errorf("Command not registered")
    }

    return m[cmd].Exec(args[2:])
}

func printHelp() {
    fmt.Println("Print big help!!!")
    for _, cmd := range m {
        cmd.PrintHelp()
    }
}

