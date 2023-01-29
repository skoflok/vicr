package explorer

import "errors"

var (
	NotSupportedLang     = errors.New("Language not supported")
	PackageFileNotExists = errors.New("Package file not found")
	BadVersion           = errors.New("Bad package version")
)
