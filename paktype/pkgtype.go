package bbpak_paktype

import "github.com/isbm/go-deb"

type PackageMeta struct {
	name     string
	version  string
	physical *deb.PackageFile
}

// Constructor
func NewPackageMeta() *PackageMeta {
	pm := new(PackageMeta)
	return pm
}

// SetVersion of the package
func (pm *PackageMeta) SetVersion(version string) {
	pm.version = version
}

// SetName of the package
func (pm *PackageMeta) SetName(name string) {
	pm.name = name
}

// Version of the package (initially assumed from the status file)
func (pm *PackageMeta) Version() string {
	return pm.version
}

// Name of the package
func (pm *PackageMeta) Name() string {
	return pm.name
}

// SetPackageFile from parsed .deb or .ipk package
func (pm *PackageMeta) SetPackageFile(pf *deb.PackageFile) {
	pm.physical = pf
}

// GetPackage returns all the information set about physical Debian package (.deb, Opkg, Ipkg)
func (pm *PackageMeta) GetPackage() *deb.PackageFile {
	return pm.physical
}
