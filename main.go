package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	confPath = "~/.nvc.yaml"
)

type Config struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

func main() {
	cmd, args, err := readConf()
	if err != nil {
		fmt.Println(err)
	}
	if err := openEditor(cmd, args...); err != nil {
		fmt.Fprint(os.Stdout, fmt.Sprintf("failed open text editor. %s\n", err.Error()))
		return
	}
}

func readConf() (string, []string, error) {
	homePath := os.Getenv("HOME")
	confPath := filepath.Join(homePath, "./.nvc.yaml")
	buf, err := ioutil.ReadFile(confPath)
	if err == nil {
		data := Config{}
		if err := yaml.Unmarshal(buf, &data); err == nil {
			return data.Command, data.Args, nil
		}
	}
	return "nvim", []string{"~/.config/nvim"}, err
}

func openEditor(cmd string, arg ...string) error {

	c := exec.Command(cmd, arg...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
