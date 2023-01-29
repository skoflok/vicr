package main

import (
	"fmt"
	"log"

	ex "github.com/skoflok/vicr/explorer"
	gu "github.com/skoflok/vicr/gitutils"
)

func main() {
	if gu.IsInstalled() == false {
		log.Fatal("Please install git\n")
	}

	cTag := gu.CurrentTag()

	fmt.Printf("Current Tag: %s\n", cTag)

	project, err := ex.NewProject("php")

	if err != nil {
		log.Fatal(err)
	}

	v, err := ex.NewVersion(cTag)
	if err != nil {
		log.Fatal(err)
	}

	ok, err := ex.ChangeVersionInProjectFile(project, v)

	fmt.Println(ok)
	fmt.Println(err)

}

func currentTag() {
	fmt.Printf("1)%s", gu.CurrentTag())
}
