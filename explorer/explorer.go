package explorer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ChangeVersionInProjectFile(p *projectType, v *version) (ok bool, err error) {
	dir, err := os.Getwd()

	if err != nil {
		return false, err
	}

	fp, err := filepath.Abs(filepath.Join(dir, p.PackageFile()))

	if err != nil {
		return false, err
	}

	b, err := os.ReadFile(fp)

	if err != nil {
		return false, err
	}

	type Schema struct {
		Version string `json:"version"`
	}

	schema := Schema{}

	if err = json.Unmarshal(b, &schema); err != nil {
		return false, err
	}

	re := regexp.MustCompile(`"version":\s{0,}("v?\d+(\.\d+)?(\.\d+)?")`)

	if re.MatchString(string(b)) == false {
		return false, NodFoundVersion
	}

	fVer, err := p.VersionFormat(v)

	if err != nil {
		return false, err
	}

	newContent := re.ReplaceAllString(string(b), fVer)

	//fmt.Println(schema)

	fmt.Println(newContent)

	return true, nil

}
