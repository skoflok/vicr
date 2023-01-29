package explorer

import "errors"

var (
	NotSupportedLang     = errors.New("Language not supported")
	NotSupportedManager  = errors.New("Package manager not supported")
	PackageFileNotExists = errors.New("Package file not found")
	BadVersion           = errors.New("Bad package version")
	NodFoundVersion      = errors.New("Version property not found in package file")
)
