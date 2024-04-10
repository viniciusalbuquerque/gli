package main

import (
	"fmt"
	"os"

	"github.com/viniciusalbuquerque/gli/src/gli"
)

// Example 1
func downloadFn(params []string) error {
    fmt.Println("Downloading file...", params)
    if len(params) == 0 {
        return fmt.Errorf("Error downloading file")
    }
    return nil
}

// Example 2
func uploadFn(params []string) error {
    fmt.Println("Uploading file...")
    if len(params) == 0 {
        return fmt.Errorf("Error uploading file")
    }
    return nil
}

func main() {

    gli.RegisterCommand("download", "Download a file from drive", downloadFn)
    gli.RegisterCommand("upload", "Upload a file to the drive", uploadFn)

    gli.ExecCommand(os.Args)
}
