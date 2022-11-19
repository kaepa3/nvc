package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if err := openEditor("nvim", "~/.config/nvim"); err != nil {
		fmt.Fprint(os.Stdout, fmt.Sprintf("failed open text editor. %s\n", err.Error()))
		return
	}
}
func openEditor(cmd string, arg string) error {
	c := exec.Command(cmd, arg)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
