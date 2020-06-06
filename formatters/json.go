package bbpak_formatters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakJSONFormat struct {
	packages map[string]*bbpak_paktype.PackageMeta
	BBPakFormatterUtils
}

func NewBBPakJSONFormat() *BBPakJSONFormat {
	bbp := new(BBPakJSONFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the JSON format (useful for integrations)
func (bbp *BBPakJSONFormat) Format() string {
	var out map[string]interface{} = make(map[string]interface{})
	for pkgName := range bbp.packages {
		pkg := bbp.packages[pkgName]
		pkgInfo := make(map[string]interface{})
		pkgInfo["license"] = pkg.GetPackage().ControlFile().Licence()
		pkgInfo["version"] = pkg.GetPackage().ControlFile().Version()
		pkgInfo["depends"] = pkg.GetPackage().ControlFile().Depends()
		pkgInfo["description"] = pkg.GetPackage().ControlFile().Description()
		pkgInfo["summary"] = pkg.GetPackage().ControlFile().Summary()

		provides := pkg.GetPackage().ControlFile().Provides()
		if provides != nil {
			pkgInfo["provides"] = provides
		}

		pkgInfo["maintainer"] = pkg.GetPackage().ControlFile().Maintainer()
		pkgInfo["pkg-size"] = bbp.ByteIEC(pkg.GetPackage().FileSize())

		pkgInfo["arch"] = pkg.GetPackage().ControlFile().Architecture()
		pkgInfo["multi-arch"] = pkg.GetPackage().ControlFile().MultiArch()

		pkgChecksum := make(map[string]string)
		pkgChecksum["md5"] = pkg.GetPackage().GetPackageChecksum().MD5()
		pkgChecksum["sha1"] = pkg.GetPackage().GetPackageChecksum().SHA1()
		pkgChecksum["sha256"] = pkg.GetPackage().GetPackageChecksum().SHA256()
		pkgInfo["checksum"] = pkgChecksum

		out[pkgName] = pkgInfo
	}

	jsonData, err := json.Marshal(out)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	var formatted bytes.Buffer
	if err := json.Indent(&formatted, jsonData, "", "  "); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	return formatted.String()
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakJSONFormat) SetPackages(packages map[string]*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
