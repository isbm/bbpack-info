package bbpak_formatters

import (
	"fmt"
	"strings"

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
	var out strings.Builder

	out.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", "Num", "Name", "Version", "Licence", "Size", "Maintainer"))
	for idx, pkgName := range bbp.pkgNameIndex {
		pkg := bbp.packages[pkgName]
		out.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s,%s\n", idx+1,
			pkg.GetPackage().ControlFile().Package(),
			pkg.GetPackage().ControlFile().Version(),
			pkg.GetPackage().ControlFile().Licence(),
			bbp.ByteIEC(pkg.GetPackage().FileSize()),
			pkg.GetPackage().ControlFile().Maintainer()))
	}
	return out.String()
}
