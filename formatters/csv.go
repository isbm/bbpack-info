package bbpak_formatters

import (
	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakCSVFormat struct {
	packages []*bbpak_paktype.PackageMeta
}

func NewBBPakCSVFormat() *BBPakCSVFormat {
	bbp := new(BBPakCSVFormat)
	bbp.packages = make([]*bbpak_paktype.PackageMeta, 0)

	return bbp
}

// Format the output to the CSV format (useful for managers :-) )
func (bbp *BBPakCSVFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakCSVFormat) SetPackages(packages []*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
