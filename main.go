package main

import (
    "errors"

    command "github.com/viniciusalbuquerque/gli/src"
)

func fn(params ...any) error {
    return errors.New("The error")
}

func main() {
    c := command.New("download", "Download a file from driver", fn)
    c.Exec(1)
}
