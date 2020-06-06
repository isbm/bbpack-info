package bbpak_formatters

import (
	"fmt"
	"strings"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakMarkdownFormat struct {
	BBPakFormatterUtils
}

func NewBBPakMarkdownFormat() *BBPakMarkdownFormat {
	bbp := new(BBPakMarkdownFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the Markdown table (useful for GitHub/-Lab Wikis)
func (bbp *BBPakMarkdownFormat) Format() string {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s|%s|\n", "Num", "Name", "Version", "Licence", "Size", "Maintainer"))
	out.WriteString("|-|-|-|-|-|-|\n")
	for idx, pkgName := range bbp.pkgNameIndex {
		pkg := bbp.packages[pkgName]
		out.WriteString(fmt.Sprintf("|%d|%s|%s|%s|%s|%s|\n", idx+1,
			pkg.GetPackage().ControlFile().Package(),
			pkg.GetPackage().ControlFile().Version(),
			pkg.GetPackage().ControlFile().Licence(),
			bbp.ByteIEC(pkg.GetPackage().FileSize()),
			pkg.GetPackage().ControlFile().Maintainer()))
	}
	return out.String()
}
