package bbpak_formatters

import bbpak "github.com/isbm/bbpack-info"

type BBPakTextFormat struct {
	packages []*bbpak.PackageMeta
}

func NewBBPakTextFormat() *BBPakTextFormat {
	bbp := new(BBPakTextFormat)
	bbp.packages = make([]*bbpak.PackageMeta, 0)

	return bbp
}

// Format the output to the ASCII text (for CLI for example)
func (bbp *BBPakTextFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakTextFormat) SetPackages(packages []*bbpak.PackageMeta) {
	bbp.packages = packages
}
