package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pu "github.com/manifoldco/promptui"
	ex "github.com/skoflok/vicr/explorer"
	gu "github.com/skoflok/vicr/gitutils"
)

var messageFlag string

func init() {
	flag.StringVar(&messageFlag, "message", "", "Commit/tag message")
}

func main() {

	flag.Parse()
	if gu.IsInstalled() == false {
		log.Fatal("Please install git\n")
	}

	if flag.NArg() == 0 {
		log.Fatal("Please specify command")
	}

	mainCmd := flag.Args()[0]

	switch mainCmd {
	case "incr":
		increaseVersion()
	case "i":
		increaseVersion()
	case "incr-commit":
	case "ic":
		increaseCommit()
	case "incr-commit-tag":
	case "ict":
		increaseCommitTag()
	case "incr-commit-tag-push":
	case "ictp":
		increaseCommitTagPush()
	default:
		log.Fatal("Command not supported")
	}

}

func increaseCommit() (version, message string) {
	version = increaseVersion()
	message = fmt.Sprintf("Release: %s. %s", version, messageFlag)
	commit(message)
	return
}

func increaseCommitTag() (version, message string) {
	version, message = increaseCommit()
	tag(version, message)
	return
}

func increaseCommitTagPush() (version, message string) {
	version, message = increaseCommitTag()
	pushCommitAndTag(version)
	return
}

func commit(message string) {
	gu.CreateCommit(message, os.Stdout)
}

func pushCommitAndTag(version string) {
	gu.PushCurrent(os.Stdout)
	gu.PushTag(version, "origin", os.Stdout)
}

func tag(version, message string) {
	gu.CreateNewTag(version, message, os.Stdout)
}

func currentTag() {
	fmt.Printf("1)%s", gu.CurrentTag())
}

func increaseVersion() string {
	cTag := gu.CurrentTag()

	fmt.Printf("Current Version: %s\n", cTag)

	project, err := ex.NewProjectType("composer")

	if err != nil {
		log.Fatal(err)
	}

	ver, err := ex.NewVersion(cTag)
	if err != nil {
		log.Fatal(err)
	}
	possibles := ver.PossibleIncreasesAsStrings()
	prompt := pu.Select{
		Label: "Choice next version:\n",
		Items: []string{
			fmt.Sprintf("Major) %s", possibles[0]),
			fmt.Sprintf("Minor) %s", possibles[1]),
			fmt.Sprintf("Patch) %s", possibles[2]),
		},
	}

	selected, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	newVer, err := ex.NewVersion(possibles[selected])

	if err != nil {
		log.Fatal(err)
	}

	ok, err := ex.ChangeVersionInProjectFile(project, newVer)
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Printf("Change Version to: %s\n", newVer)
	} else {
		log.Fatal("Something went wrong\n")
	}

	return newVer.String()

}
