package gitutils

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

func IsInstalled() bool {
	if _, err := exec.LookPath("git"); err != nil {
		return false
	}
	return true
}

func CurrentTag() string {
	s := "tag --sort=-creatordate"
	cmd := exec.Command("git", strings.Split(s, " ")...)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running command 'git tag args...' : %v", err)
	}

	cmdHead := exec.Command("head", "-1")

	stdin, err := cmdHead.StdinPipe()

	if err != nil {
		log.Fatalf("Error stdin 'head args...' command %v", err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, string(output))
	}()

	out, err := cmdHead.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}
