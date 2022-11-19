package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	launchVim()
}
func launchVim() int {
	// Open text editor
	err := openEditor("nvim")
	if err != nil {
		fmt.Fprint(os.Stdout, fmt.Sprintf("failed open text editor. %s\n", err.Error()))
		return 1
	}
	return 0
}

func openEditor(program string) error {
	c := exec.Command(program, "~/.config/nvim")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
