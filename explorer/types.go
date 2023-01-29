package explorer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type project struct {
	lang        string
	packageFile string
}

type version struct {
	ver                 string
	major, minor, patch int
}

var typeDictionary = map[string]string{
	"php":        "composer.json",
	"javascript": "package.json",
	"js":         "package.json",
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

	v.ver = ver

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

func NewProject(lang string) (p *project, err error) {
	if pfile, ok := typeDictionary[lang]; ok == true {
		return &project{lang, pfile}, nil
	}

	err = fmt.Errorf("Input %s. %v", lang, NotSupportedLang)

	return p, err
}

func (p *project) PackageFile() string {
	return p.packageFile
}

func (p *project) Language() string {
	return p.lang
}

func (p *project) VersionFormat(v *version) (string, error) {
	switch p.Language() {
	case "php":
		return fmt.Sprintf("\"version\": \"%s\"", v), nil
	case "javascript":
		return fmt.Sprintf("\"version\": \"%s\"", v), nil
	case "js":
		return fmt.Sprintf("\"version\": \"%s\"", v), nil
	default:
		return "", NotSupportedLang
	}

	return "", NotSupportedLang
}
