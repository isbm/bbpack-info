package bbpak_formatters

import (
	"fmt"
	"sort"
	"strings"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakSTDOUTFormat struct {
	BBPakFormatterUtils
}

func NewBBPakSTDOUTFormat() *BBPakSTDOUTFormat {
	bbp := new(BBPakSTDOUTFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the ASCII text (for CLI for example)
func (bbp *BBPakSTDOUTFormat) Format() string {
	var out strings.Builder

	for idx, pkgName := range bbp.pkgNameIndex {
		pkg := bbp.packages[pkgName]
		out.WriteString(fmt.Sprintf("%3d. %s\n", idx+1, pkg.GetPackage().ControlFile().Package()))

		deps := pkg.GetPackage().ControlFile().Depends()
		if len(deps) > 0 {
			sort.Strings(deps)
			if len(deps) > 2 {
				out.WriteString("     Depends:\n")
				for didx, f := range deps {
					out.WriteString(fmt.Sprintf("     %3d. %s\n", didx+1, f))
				}
			} else {
				out.WriteString(fmt.Sprintf("     Depends:  %s\n", deps[0]))
			}
		}

		provides := pkg.GetPackage().ControlFile().Provides()
		if len(provides) > 0 {
			sort.Strings(provides)
			if len(provides) > 2 {
				out.WriteString("     Provides:\n")
				for didx, f := range provides {
					out.WriteString(fmt.Sprintf("     %3d. %s\n", didx+1, f))
				}
			} else {
				out.WriteString(fmt.Sprintf("     Provides: %s\n", provides[0]))
			}
		}

		out.WriteString(fmt.Sprintf("\n     License: %s, Maintainer: %s, Size: %s\n",
			pkg.GetPackage().ControlFile().Licence(),
			pkg.GetPackage().ControlFile().Maintainer(),
			bbp.ByteIEC(pkg.GetPackage().FileSize())))
		out.WriteString("\n")
	}

	return out.String()
}
