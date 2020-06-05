package bbpak

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/isbm/go-deb"
)

type PackageMeta struct {
	name    string
	version string
}

// BBPakMatcher class
type BBPakMatcher struct {
	root     string
	manifest string
	pkgs     []PackageMeta
	pkgPaths []string
}

// Costructor
func NewBBPakMatcher(path string) *BBPakMatcher {
	bb := new(BBPakMatcher)
	bb.root = path
	bb.pkgs = make([]PackageMeta, 0)
	bb.pkgPaths = make([]string, 0)
	return bb
}

func (bb *BBPakMatcher) parsePackageSection(buff []string) *PackageMeta {
	meta := new(PackageMeta)
	for _, line := range buff {
		parts := strings.Split(line, " ")
		if len(parts) < 1 {
			continue
		}
		switch parts[0] {
		case "Package:":
			meta.name = parts[1]
		case "Version:":
			meta.version = parts[1]
		}
	}
	if meta.name == "" {
		panic("Oops, name of the package is missing. The rest of the data does not matter therefore. Your status file seems just broken.")
	}

	return meta
}

func (bb *BBPakMatcher) ParseManifestPackages() {
	f, err := os.Open(bb.manifest)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var buff []string = nil
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if buff == nil {
			buff = make([]string, 0)
		}

		if line == "" {
			bb.pkgs = append(bb.pkgs, *bb.parsePackageSection(buff))
			buff = nil
		} else {
			buff = append(buff, line)
		}
	}
	if buff != nil {
		bb.pkgs = append(bb.pkgs, *bb.parsePackageSection(buff))
	}
}

func (bb *BBPakMatcher) findManifest(pth string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.HasSuffix(pth, ".rootfs.opkg.status") {
		bb.manifest = pth
	} else if strings.HasSuffix(pth, ".ipk") {
		bb.pkgPaths = append(bb.pkgPaths, pth)
	}
	return nil
}

func (bb *BBPakMatcher) FindPhysicalPackages() {
	// This is terrible
	for _, pkg := range bb.pkgs {
		for _, pth := range bb.pkgPaths {
			if strings.HasPrefix(path.Base(pth), pkg.name+"_"+pkg.version) {
				p, err := deb.OpenPackageFile(pth, false)
				if err != nil {
					fmt.Println("Error opening package:", err.Error())
				}
				fmt.Println("Package:", p.ControlFile().Package())
				fmt.Println("Maintainer:", p.ControlFile().Maintainer())
				fmt.Println("Version:", p.ControlFile().Version())
				os.Exit(1)
			}
		}
	}
}

func (bb *BBPakMatcher) FindManifests() {
	deployed := path.Join(bb.root, "build", "tmp", "deploy")
	err := filepath.Walk(deployed, bb.findManifest)
	if err != nil && bb.manifest == "" {
		fmt.Println(">>>", err)
	} else {
		fmt.Println("Manifest:", bb.manifest)

		bb.ParseManifestPackages()
		bb.FindPhysicalPackages()

		for num, p := range bb.pkgs {
			fmt.Printf("%d. %s (%s)\n", num+1, p.name, p.version)
		}
	}
}
