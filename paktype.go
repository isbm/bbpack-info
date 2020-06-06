package bbpak

import "github.com/isbm/go-deb"

type PackageMeta struct {
	name     string
	version  string
	physical *deb.PackageFile
}

// Constructor
func NewPackageMeta(name string) *PackageMeta {
	pm := new(PackageMeta)
	pm.name = name
	return pm
}

// SetPackageFile from parsed .deb or .ipk package
func (pm *PackageMeta) SetPackageFile(pf *deb.PackageFile) {
	pm.physical = pf
}
