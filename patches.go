package bbpak

/*
	Patch tracker is to find actually applied patches per a package,
	regardless what is actually placed in meta and/or defined anywhere (any recipes etc).
*/

type BBPakPatchesTracker struct {
	pkgName string
}

func NewBBPakPatchesTracker(pkgName string) *BBPakPatchesTracker {
	bbpt := new(BBPakPatchesTracker)
	bbpt.pkgName = pkgName
	return bbpt
}

// GetAllPatches which are defined to the package (not the fact they are actually applied).
func (bbpt *BBPakPatchesTracker) GetAllPatches() {
}

// GetAppliedPatches which are actually are used in the given patch and compiled with.
func (bbpt *BBPakPatchesTracker) GetAppliedPatches() {
}

// BuildChangelog to the package.
func (bbpt *BBPakPatchesTracker) BuildChangelog() {
	// Cache changelog to the package and remove if rootfs.status.opkg has been updated (timestamp)
}
