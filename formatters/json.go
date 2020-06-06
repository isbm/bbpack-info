package bbpak_formatters

import bbpak_paktype "github.com/isbm/bbpack-info/paktype"

type BBPakJSONFormat struct {
	packages []*bbpak_paktype.PackageMeta
}

func NewBBPakJSONFormat() *BBPakCSVFormat {
	bbp := new(BBPakCSVFormat)
	bbp.packages = make([]*bbpak_paktype.PackageMeta, 0)

	return bbp
}

// Format the output to the JSON format (useful for integrations)
func (bbp *BBPakJSONFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakJSONFormat) SetPackages(packages []*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
