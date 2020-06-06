package bbpak_formatters

import (
	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakMarkdownFormat struct {
	packages map[string]*bbpak_paktype.PackageMeta
	BBPakFormatterUtils
}

func NewBBPakMarkdownFormat() *BBPakMarkdownFormat {
	bbp := new(BBPakMarkdownFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the Markdown table (useful for GitHub/-Lab Wikis)
func (bbp *BBPakMarkdownFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakMarkdownFormat) SetPackages(packages map[string]*bbpak_paktype.PackageMeta) {
	bbp.packages = packages
}
