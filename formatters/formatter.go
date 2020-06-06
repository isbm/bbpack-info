package bbpak_formatters

import (
	bbpak "github.com/isbm/bbpack-info"
)

type BBPakFormatter interface {
	Format() string
	SetPackages(packages []*bbpak.PackageMeta)
}
