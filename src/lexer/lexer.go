package lexer

import "fmt"

type TokenType int

const (
	ID TokenType = iota
	Option
)

type Token struct {
	Value string
	Kind TokenType
}


func newOption(param string) *Token {
	return &Token {
		Value: param,
		Kind: Option,
	}
}

func parseOption(param string) (error, []*Token) {
	if len(param) == 1 {
		return fmt.Errorf("Error when parsing option"), nil
	}

	var opts []*Token
	switch(param[1]) {
	case '-':
		opt := newOption(param[2:])
		opts = append(opts, opt)
	default:
	for i := 1; i < len(param); i++ {
		opt := newOption(string(param[i]))
		opts = append(opts, opt)
		}
	}

	return nil, opts
}

func parse(param string) *Token {
	return &Token {
		Value: param,
		Kind: ID,
	}
}

func Parse(param string) (error, []*Token) {
	if len(param) == 0 {
		return fmt.Errorf("Error Parse"), nil
	}

	switch param[0] {
	case '-':
		return parseOption(param)
	default:
		token := parse(param)
		tkns := [1]*Token{token}
		return nil, tkns[:]

	}
}
