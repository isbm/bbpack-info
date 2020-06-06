package bbpak_formatters

import (
	"fmt"
	"sort"
	"strings"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakTextFormat struct {
	packages     map[string]*bbpak_paktype.PackageMeta
	pkgNameIndex []string
	BBPakFormatterUtils
}

func NewBBPakTextFormat() *BBPakTextFormat {
	bbp := new(BBPakTextFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

func (bbp *BBPakTextFormat) genIndex() {
	if bbp.pkgNameIndex == nil {
		bbp.pkgNameIndex = make([]string, 0)
		for pkName := range bbp.packages {
			bbp.pkgNameIndex = append(bbp.pkgNameIndex, pkName)
		}
		sort.Strings(bbp.pkgNameIndex)
	}
}

// Format the output to the ASCII text (for CLI for example)
func (bbp *BBPakTextFormat) Format() string {
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

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakTextFormat) SetPackages(packages map[string]*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
	bbp.pkgNameIndex = nil
	bbp.genIndex()
}
