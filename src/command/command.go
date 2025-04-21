package command

import (
	"fmt"

	"github.com/viniciusalbuquerque/gli/src/lexer"
)

type Command struct {
    Name string
    Description string
    Subcommands []*Command
    Fn func([]string, []*Option) error
    Opts []*Option
}

type Option struct {
    name string
    prefix string
    descrition string
}

func NewOption(name string, prefix string, description string) *Option {
    return &Option{
        name: name,
        prefix: prefix,
        descrition: description,
    }
}

func New(name string, desc string, fn func([]string, []*Option) error, opts []*Option) *Command {
    return &Command{
        Name: name,
        Description: desc,
        Subcommands: nil,
        Fn: fn,
        Opts: opts,
    }
}

func (cmd *Command) AddOpt(opt *Option) {
    cmd.Opts = append(cmd.Opts, opt)
}

func (cmd *Command) AddSubcommand(subcmd *Command) {
    cmd.Subcommands = append(cmd.Subcommands, subcmd)
}


func (cmd Command) Exec(params []string) error {
    err := cmd.Fn(params, nil)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}


func (cmd Command) Execute(tokens []*lexer.Token, idx int) error {
    subcmds := cmd.Subcommands
    if subcmds != nil && len(subcmds) > 0 {
        for i := idx; i < len(tokens); i++ {
            token := tokens[i]
            for j := range subcmds {
                if token.Kind == lexer.ID && token.Value == subcmds[j].Name {
                    return subcmds[j].Execute(tokens, i)
                }
            }
        }
    }

    hasOptToken := false
    var opts []*Option
    for _, token := range tokens {
        if token.Kind == lexer.Option {
            hasOptToken = true
            for _, opt := range cmd.Opts {
                if token.Value == opt.name || token.Value == opt.prefix {
                    opts = append(opts, opt)
                }
            }
        }
    }

    if hasOptToken && (opts == nil || len(opts) == 0) {
        return fmt.Errorf("Invalid option for cmd: %s", cmd.Name)
    }

    var params []string
    for i := idx; i < len(tokens); i++ {
        token := tokens[i]
        if token.Kind == lexer.ID && token.Value != cmd.Name {
            params = append(params, token.Value)
        }
    }

    return cmd.Fn(params, opts)
}

func (cmd Command) PrintHelp() {
    fmt.Println(cmd.Name)
    fmt.Println(cmd.Description)
    fmt.Printf("\n\tOptions\n\n")
    for i := range cmd.Opts {
        opt := cmd.Opts[i]
        fmt.Printf("\t--%s -%s\n", opt.name, opt.prefix)
        fmt.Printf("\t%s\n\n", opt.descrition)
    }
}

