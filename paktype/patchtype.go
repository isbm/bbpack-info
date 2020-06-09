package bbpak_paktype

import (
	"fmt"
	"os"
	"path"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

type BBPakPatch struct {
	fptr          os.FileInfo
	affectedFiles []*gitdiff.File
	header        *gitdiff.PatchHeader
	origPath      string
	preamble      string
	name          string
}

func NewBBPakPatch() *BBPakPatch {
	return new(BBPakPatch)
}

func (bp *BBPakPatch) LoadPatch(pth string) error {
	var err error
	bp.origPath = pth
	bp.fptr, err = os.Stat(pth)
	if os.IsNotExist(err) {
		return fmt.Errorf("File %s does not exist", pth)
	} else if os.IsPermission(err) {
		return fmt.Errorf("Access denied to look into %s", pth)
	}
	if err != nil {
		return err // Generic
	}

	patch, err := os.Open(bp.origPath)
	if err != nil {
		return err
	}

	bp.affectedFiles, bp.preamble, err = gitdiff.Parse(patch)
	if err != nil {
		return err
	}

	bp.header, err = gitdiff.ParsePatchHeader(bp.preamble)
	if err != nil {
		return err
	}

	bp.name = path.Base(bp.fptr.Name())

	return nil
}

// GetPatchPreamble returns patch preamble as a string
func (bp *BBPakPatch) GetPatchPreamble() string {
	return bp.preamble
}

// GetPatchHeader returns patch header
func (bp *BBPakPatch) GetPatchHeader() *gitdiff.PatchHeader {
	return bp.header
}

// GetAffectedFiles returns files that patch is applied to
func (bp *BBPakPatch) GetAffectedFiles() []*gitdiff.File {
	return bp.affectedFiles
}

// GetChangelogEntry returns a changelog entry why this patch was added.
func (bp *BBPakPatch) GetChangelogEntry() string {
	// This parses the patch itself and also links to the git source of the package inside the Yocto
	// If Git entry is not found, raw patch data is returned
	return ""
}

// GetCommitChecksum of the patch from the source tree
func (bp *BBPakPatch) GetCommitChecksum() string {
	return ""
}
