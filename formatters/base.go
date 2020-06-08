package bbpak_formatters

import (
	"fmt"
	"sort"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakFormatterUtils struct {
	packages     map[string]*bbpak_paktype.PackageMeta
	pkgNameIndex []string
}

func (bbp *BBPakFormatterUtils) genIndex() {
	if bbp.pkgNameIndex == nil {
		bbp.pkgNameIndex = make([]string, 0)
		for pkName := range bbp.packages {
			bbp.pkgNameIndex = append(bbp.pkgNameIndex, pkName)
		}
		sort.Strings(bbp.pkgNameIndex)
	}
}

// IsSummary of all packages (true) or not
func (bbp *BBPakFormatterUtils) IsSummary() bool {
	return len(bbp.packages) != 1
}

func (bbp *BBPakFormatterUtils) ByteIEC(b uint64) string {
	const unit = 0x400
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakFormatterUtils) SetPackages(packages map[string]*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
	bbp.pkgNameIndex = nil
	bbp.genIndex()
}
