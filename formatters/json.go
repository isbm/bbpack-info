package bbpak_formatters

import bbpak "github.com/isbm/bbpack-info"

type BBPakJSONFormat struct {
	packages []*bbpak.PackageMeta
}

func NewBBPakJSONFormat() *BBPakCSVFormat {
	bbp := new(BBPakCSVFormat)
	bbp.packages = make([]*bbpak.PackageMeta, 0)

	return bbp
}

// Format the output to the JSON format (useful for integrations)
func (bbp *BBPakJSONFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakJSONFormat) SetPackages(packages []*bbpak.PackageMeta) {
	bbp.packages = packages
}
