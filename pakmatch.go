package bbpak

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	bbpak_formatters "github.com/isbm/bbpack-info/formatters"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
	"github.com/isbm/go-deb"
)

// BBPakMatcher class
type BBPakMatcher struct {
	root           string
	manifestTarget string
	manifest       string // The Chosen One TM
	manifests      []string
	pkgs           map[string]*bbpak_paktype.PackageMeta
	pkgPaths       []string
}

// Costructor
func NewBBPakMatcher(path string) *BBPakMatcher {
	bb := new(BBPakMatcher)
	bb.root = path
	bb.pkgs = make(map[string]*bbpak_paktype.PackageMeta)
	bb.pkgPaths = make([]string, 0)
	bb.manifests = make([]string, 0)
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
	var pkgMeta *bbpak_paktype.PackageMeta
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if buff == nil {
			buff = make([]string, 0)
		}

		if line == "" {
			pkgMeta = bb.parsePackageSection(buff)
			bb.pkgs[pkgMeta.Name()] = pkgMeta
			buff = nil
		} else {
			buff = append(buff, line)
		}
	}
	if buff != nil {
		pkgMeta = bb.parsePackageSection(buff)
		bb.pkgs[pkgMeta.Name()] = pkgMeta
	}
}

func (bb *BBPakMatcher) findManifest(pth string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.HasSuffix(pth, ".rootfs.opkg.status") {
		if strings.Contains(pth, bb.manifestTarget) && bb.manifest == "" {
			bb.manifest = pth
		}
		bb.manifests = append(bb.manifests, pth)
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

func (bb *BBPakMatcher) Format(fmtype string) {
	var formatter bbpak_formatters.BBPakFormatter
	switch fmtype {
	case "csv":
		formatter = bbpak_formatters.NewBBPakCSVFormat()
	default:
		formatter = bbpak_formatters.NewBBPakTextFormat()
	}
	formatter.SetPackages(bb.pkgs)
	fmt.Println(formatter.Format())
}

// SetTargetManifest sets a manifest search criteria that will eventually turn it into a full path, if found.
func (bb *BBPakMatcher) SetTargetManifest(target string) {
	bb.manifestTarget = target
}

func (bb *BBPakMatcher) FindManifests() ([]string, error) {
	deployed := path.Join(bb.root, "build", "tmp", "deploy")
	err := filepath.Walk(deployed, bb.findManifest)
	if err == nil && bb.manifest == "" {
		err = fmt.Errorf("No manifest has been found")
	}
	return bb.manifests, err
}
