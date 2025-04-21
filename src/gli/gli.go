package gli

import (
	"fmt"

	"github.com/viniciusalbuquerque/gli/src/command"
	"github.com/viniciusalbuquerque/gli/src/lexer"
)

var m map[string]*command.Command
var commands []*command.Command

func RegisterCommand2(cmd *command.Command) {
    commands = append(commands, cmd)
}

func RegisterCommand(name string, desc string, fn func([]string, []*command.Option) error, opts []command.Option) {

    if m == nil {
        m = make(map[string]*command.Command)
    }

    m[name] = command.New(name, desc, fn, nil)
}

func execCommand(tokens []*lexer.Token) error {
    for idx := range tokens {
        token := tokens[idx]
        for j := range commands {
            if token.Kind == lexer.ID && token.Value == commands[j].Name {
                return commands[j].Execute(tokens, 0)
            }
        }
    }
    return fmt.Errorf("Command not registered")
}

func ExecCommand2(args []string) error {
    if len(args) < 2  {
        printHelp()
        return fmt.Errorf("Not enough arguments.")
    }

    if commands == nil || len(commands) == 0 {
        return fmt.Errorf("No commands were registered")
    }

    var tokens []*lexer.Token
    for i := 1; i < len(args); i++ {
        err, tks := lexer.Parse(args[i])
        if err != nil {
            return fmt.Errorf("Exec command parse error")
        }

        for idx := range tks {
            tokens = append(tokens, tks[idx])
        }
    }

    return execCommand(tokens)
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

    var tokens []*lexer.Token
    for i := 2; i < len(args); i++ {
        err, tks := lexer.Parse(args[i])
        if err != nil {
            return fmt.Errorf("Exec command parse error")
        }
        for idx := range tks {
            tokens = append(tokens, tks[idx])
        }
    }

    for idx := range tokens {
        fmt.Println(tokens[idx])
    }

    return m[cmd].Exec(args[2:])
}

func PrintHelp() {
    for i := range commands {
        commands[i].PrintHelp()
        fmt.Println()
    }
}

//TODO: Format print correctly
func printHelp() {
    fmt.Println("Print big help!!!")
    for _, cmd := range m {
        cmd.PrintHelp()
    }
}

