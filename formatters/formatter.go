package bbpak_formatters

import (
	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakFormatter interface {
	Format() string
	SetPackages(packages map[string]*bbpak_paktype.PackageMeta)
	ByteIEC(b uint64) string
}
