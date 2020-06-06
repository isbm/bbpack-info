package bbpak_formatters

import (
	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakCSVFormat struct {
	BBPakFormatterUtils
}

func NewBBPakCSVFormat() *BBPakCSVFormat {
	bbp := new(BBPakCSVFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the CSV format (useful for managers :-) )
func (bbp *BBPakCSVFormat) Format() string {
	return "csv is not yet implemented"
}
