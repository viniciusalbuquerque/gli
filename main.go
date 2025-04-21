package main

import (
	"fmt"
	"os"

	"github.com/viniciusalbuquerque/gli/src/command"
	"github.com/viniciusalbuquerque/gli/src/gli"
)

// Example 1
func downloadFn(params []string, opts []*command.Option) error {
    fmt.Println("Downloading file...", params, opts[0])
    if len(params) == 0 {
        return fmt.Errorf("Error downloading file")
    }
    return nil
}

// Example 2
func uploadFn(params []string, opts []*command.Option) error {
    fmt.Println("Uploading file...", params, opts)
    if len(params) == 0 {
        return fmt.Errorf("Error uploading file")
    }
    return nil
}

func main() {

    opt1 := command.NewOption("test", "t", "This is a test option")
    opt2 := command.NewOption("test2", "r", "This is a test option 2")
    var opts []*command.Option
    opts = append(opts, opt1)
    opts = append(opts, opt2)

    cmd1 := command.New("download", "Download a file from drive", downloadFn, opts)
    cmd2 := command.New("upload", "Upload a file to the drive", uploadFn, nil)

    cmd1.AddSubcommand(cmd2)
    gli.RegisterCommand2(cmd1)

    err := gli.ExecCommand2(os.Args)
    if err != nil {
	fmt.Println(err.Error())
    }
}
