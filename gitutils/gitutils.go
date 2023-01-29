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

func CreateNewTag(tag, message string, out io.Writer) {
	if message == "" {
		message = tag
	}

	cmd := exec.Command("git", "tag", "-a", tag, "-m", message)

	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running command tag create 'git tag args...' : %v", err)
	}
}

func CreateCommit(message string, out io.Writer) {
	cmd := exec.Command("git", "commit", "-am", message)

	cmd.Stdout = out

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running commit create 'git commit args...' : %v", err)
	}
}

func PushCurrent(out io.Writer) {
	cmd := exec.Command("git", "push")

	cmd.Stdout = out

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running push 'git push' : %v", err)
	}

}

func PushTag(tag, remote string, out io.Writer) {
	if remote == "" {
		remote = "origin"
	}

	cmd := exec.Command("git", "push", remote, tag)

	cmd.Stdout = out

	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error running command push new tag 'git push args...' : %v", err)
	}
}
