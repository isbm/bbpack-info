package bbpak

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
	"github.com/isbm/go-deb"
)

// BBPakMatcher class
type BBPakMatcher struct {
	root     string
	manifest string
	pkgs     []*bbpak_paktype.PackageMeta
	pkgPaths []string
}

// Costructor
func NewBBPakMatcher(path string) *BBPakMatcher {
	bb := new(BBPakMatcher)
	bb.root = path
	bb.pkgs = make([]*bbpak_paktype.PackageMeta, 0)
	bb.pkgPaths = make([]string, 0)
	return bb
}

func (bb *BBPakMatcher) parsePackageSection(buff []string) *bbpak_paktype.PackageMeta {
	meta := new(bbpak_paktype.PackageMeta)
	for _, line := range buff {
		parts := strings.Split(line, " ")
		if len(parts) < 1 {
			continue
		}
		switch parts[0] {
		case "Package:":
			meta.SetName(parts[1])
		case "Version:":
			meta.SetVersion(parts[1])
		}
	}
	if meta.Version() == "" {
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
			bb.pkgs = append(bb.pkgs, bb.parsePackageSection(buff))
			buff = nil
		} else {
			buff = append(buff, line)
		}
	}
	if buff != nil {
		bb.pkgs = append(bb.pkgs, bb.parsePackageSection(buff))
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

func (bb *BBPakMatcher) prepareVersion(ver string) string {
	if strings.Contains(ver, ":") {
		return strings.Split(ver, ":")[1]
	}
	return ver
}

func (bb *BBPakMatcher) FindPhysicalPackages() {
	// This is terrible
	missing := make([]string, 0)
	for _, pkg := range bb.pkgs {
		found := false
		for _, pth := range bb.pkgPaths {
			if strings.HasPrefix(path.Base(pth), pkg.Name()+"_") && strings.Contains(pth, bb.prepareVersion(pkg.Version())) {
				p, err := deb.OpenPackageFile(pth, false)
				if err != nil {
					fmt.Println("Error opening package:", err.Error())
				}
				pkg.SetPackageFile(p)
				found = true
			}
		}
		if !found {
			missing = append(missing, pkg.Name()+"("+pkg.Version()+")")
		}
	}

	if len(missing) > 0 {
		fmt.Println("Missing:", len(missing))
		for i, p := range missing {
			fmt.Println(i+1, p)
		}
		panic("Some packages has been missing!")
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

	}
}
