package bbpak

import (
	"fmt"
	"path"
	"sort"
	"strings"
	"sync"

	"github.com/karrick/godirwalk"
)

/*
	Patch tracker is to find actually applied patches per a package,
	regardless what is actually placed in meta and/or defined anywhere (any recipes etc).
*/

type BBPakPatchesTracker struct {
	pkgName          string
	root             string
	allPatches       map[string]interface{}
	appliedPatches   map[string]string
	applyPatchLogPth string
	mu               *sync.Mutex
}

func NewBBPakPatchesTracker(root string, pkgName string) *BBPakPatchesTracker {
	bbpt := new(BBPakPatchesTracker)
	bbpt.root = ResolveIfSymlink(root)
	bbpt.pkgName = pkgName
	bbpt.allPatches = make(map[string]interface{})
	bbpt.appliedPatches = make(map[string]string)

	bbpt.mu = new(sync.Mutex)
	return bbpt
}

// Patches that are attached to the package source
func (bbpt *BBPakPatchesTracker) filterSourcePatch(pth string, info *godirwalk.Dirent) error {
	if strings.Contains(pth, "build/tmp/") {
		return fmt.Errorf(pth)
	}
	if strings.Contains(pth, "/"+bbpt.pkgName+"/") && strings.HasSuffix(pth, ".patch") {
		bbpt.allPatches[path.Base(pth)] = nil
	}
	return nil
}

// Patches that are actually applied
func (bbpt *BBPakPatchesTracker) filterDeployedPatch(pth string, info *godirwalk.Dirent) error {
	if strings.Contains(pth, "/"+bbpt.pkgName+"/") && strings.HasSuffix(pth, ".patch") && strings.Contains(pth, "/tmp/work/") {
		bbpt.appliedPatches[path.Base(pth)] = "" // Then resolve paths and ordering by do_patch log
	}
	return nil
}

// GetAllPatches which are defined to the package (not the fact they are actually applied).
func (bbpt *BBPakPatchesTracker) GetAllPatches() []string {
	patches := make([]string, 0)
	err := godirwalk.Walk(bbpt.root, &godirwalk.Options{
		Unsorted: true,
		Callback: bbpt.filterSourcePatch,
		ErrorCallback: func(pth string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
	if err != nil {
		fmt.Println("Error getting all patches:", err.Error())
		return patches
	}

	for p := range bbpt.allPatches {
		patches = append(patches, p)
	}
	sort.Strings(patches)
	return patches
}

// GetAppliedPatches which are actually are used in the given patch and compiled with.
func (bbpt *BBPakPatchesTracker) GetAppliedPatches() []string {
	patches := make([]string, 0)
	err := godirwalk.Walk(path.Join(bbpt.root, "build", "tmp", "work"), &godirwalk.Options{
		Unsorted: true,
		Callback: bbpt.filterDeployedPatch,
		ErrorCallback: func(pth string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
	if err != nil {
		fmt.Println("Error getting all patches:", err.Error())
		return patches
	}

	for p := range bbpt.appliedPatches {
		patches = append(patches, p)
	}
	sort.Strings(patches)
	return patches
}

// BuildChangelog to the package.
func (bbpt *BBPakPatchesTracker) BuildChangelog() {
	// Cache changelog to the package and remove if rootfs.status.opkg has been updated (timestamp)
}
