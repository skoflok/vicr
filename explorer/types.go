package explorer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type version struct {
	major, minor, patch int
}

type projectType struct {
	manager  string
	language string
	file     string
}

var typeDictionary = map[string]*projectType{
	"composer": &projectType{"composer", "php", "composer.json"},
	//	"javascript": "package.json",
	//	"js":         "package.json",
}

func NewVersion(ver string) (v *version, err error) {

	s := ""
	mmp := []string{}
	re := regexp.MustCompile(`v?\d+(\.\d+)?(\.\d+)?`)
	if re.MatchString(ver) == false {
		return nil, fmt.Errorf("Bad version format. Input: %s. Error: %v", ver, BadVersion)
	}

	s = strings.Trim(ver, "v")
	mmp = strings.Split(s, ".")

	v = &version{}

	if v.major, err = strconv.Atoi(mmp[0]); err != nil {
		return nil, fmt.Errorf("Bad major number %s. Error: %v", mmp[0], BadVersion)
	}
	v.minor = 0
	v.patch = 0

	if len(mmp) > 1 {
		if v.minor, err = strconv.Atoi(mmp[1]); err != nil {
			return nil, fmt.Errorf("Bad minor number %s. Error: %v", mmp[1], BadVersion)
		}
	}
	if len(mmp) > 2 {
		if v.patch, err = strconv.Atoi(mmp[2]); err != nil {
			return nil, fmt.Errorf("Bad patch number %s. Error: %v", mmp[2], BadVersion)
		}
	}

	return v, nil
}

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func (v *version) PossibleIncreases() []*version {
	return []*version{
		&version{v.major + 1, 0, 0},
		&version{v.major, v.minor + 1, 0},
		&version{v.major, v.minor, v.patch + 1},
	}

}

func (v *version) PossibleIncreasesAsStrings() []string {
	s := []string{}

	for _, v := range v.PossibleIncreases() {
		s = append(s, v.String())
	}

	return s
}

func NewProjectType(manager string) (p *projectType, err error) {
	if ptype, ok := typeDictionary[manager]; ok == true {
		return ptype, nil
	}

	err = fmt.Errorf("Input %s. %v", manager, NotSupportedManager)

	return p, err
}

func (p *projectType) PackageFile() string {
	return p.file
}

func (p *projectType) Language() string {
	return p.language
}
func (p *projectType) Manager() string {
	return p.manager
}
func (p *projectType) VersionFormat(v *version) (string, error) {
	switch p.Manager() {
	case "composer":
		return fmt.Sprintf("\"version\": \"%s\"", v), nil
	default:
		return "", NotSupportedManager
	}

	return "", NotSupportedManager
}
