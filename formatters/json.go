package bbpak_formatters

import (
	"bytes"
	"encoding/json"

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
		out[pkgName] = pkgInfo
	}

	jsonData, err := json.Marshal(out)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	var formatted bytes.Buffer
	json.Indent(&formatted, jsonData, "", "\t")

	return formatted.String()
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakJSONFormat) SetPackages(packages map[string]*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
