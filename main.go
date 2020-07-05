package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/idtools"
)

func exitError(err interface{}) {
	fmt.Fprintf(os.Stderr, "Error: %v", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		exitError("invalid arguments")
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		exitError(err)
	}

	if resp.StatusCode != 200 {
		exitError(fmt.Sprintf("Status response was not 200; was %d (%v)", resp.StatusCode, resp.Status))
	}

	if err := os.MkdirAll(os.Args[2], os.ModePerm); err != nil { // rely on umask
		exitError(err)
	}

	options := &archive.TarOptions{
		ChownOpts: &idtools.Identity{
			UID: os.Getuid(),
			GID: os.Getgid(),
		},
		Compression: archive.Gzip,
	}

	if err := archive.Untar(resp.Body, os.Args[2], options); err != nil {
		exitError(err)
	}

	fmt.Printf("Unpacking of %q from URL %q was successful.", os.Args[2], os.Args[1])
}
