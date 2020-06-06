package bbpak_formatters

import (
	"fmt"
	"strings"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakTextFormat struct {
	packages []*bbpak_paktype.PackageMeta
}

func NewBBPakTextFormat() *BBPakTextFormat {
	bbp := new(BBPakTextFormat)
	bbp.packages = make([]*bbpak_paktype.PackageMeta, 0)

	return bbp
}

// Format the output to the ASCII text (for CLI for example)
func (bbp *BBPakTextFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakTextFormat) SetPackages(packages []*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
