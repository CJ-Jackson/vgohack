package main

import (
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 2 || os.Args[1] != "mod" {
		cmd := exec.Command("go", os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return
	}

	modFlagMap := map[string]string{
		"-require": "edit",
		"-fix":     "fix",
		"-graph":   "graph",
		"-init":    "init",
		"-sync":    "tidy",
		"-vendor":  "vendor",
		"-verify":  "verify",
	}

	arg2 := os.Args[2]
	if arg2[0] == '-' {
		arg2 = modFlagMap[arg2]
	}

	args := []string{os.Args[1], arg2}
	if len(os.Args) > 3 {
		args = append(args, os.Args[3:]...)
	}

	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
