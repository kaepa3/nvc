package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if err := openEditor("nvim"); err != nil {
		fmt.Fprint(os.Stdout, fmt.Sprintf("failed open text editor. %s\n", err.Error()))
		return
	}
}
func openEditor(program string) error {
	c := exec.Command(program, "~/.config/nvim")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
