package bbpak_formatters

import (
	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakFormatter interface {
	Format() string
	SetPackages(packages []*bbpak_paktype.PackageMeta)
}
