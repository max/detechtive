package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var packs = []string{"ruby", "nodejs", "python", "java", "php", "go"}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no directory specified (i.e. `detechtive <path>`)")
		os.Exit(1)
	}

	path := args[0]
	if err := pathIsValid(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range packs {
		cmd := exec.Command(fmt.Sprintf("./packs/%s/detect", p), path)
		if err := cmd.Run(); err == nil {
			fmt.Println(p)
			break
		}
	}
}

func pathIsValid(path string) error {
	p, err := os.Stat(path)
	if err != nil {
		return errors.New("no such directory")
	}

	if !p.IsDir() {
		return errors.New("the specified path is not a directory")
	}

	return nil
}
