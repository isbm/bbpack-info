package bbpak_formatters

import bbpak "github.com/isbm/bbpack-info"

type BBPakCSVFormat struct {
	packages []*bbpak.PackageMeta
}

func NewBBPakCSVFormat() *BBPakCSVFormat {
	bbp := new(BBPakCSVFormat)
	bbp.packages = make([]*bbpak.PackageMeta, 0)

	return bbp
}

// Format the output to the CSV format (useful for managers :-) )
func (bbp *BBPakCSVFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakCSVFormat) SetPackages(packages []*bbpak.PackageMeta) {
	bbp.packages = packages
}
