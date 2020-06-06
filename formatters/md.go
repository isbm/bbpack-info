package bbpak_formatters

import bbpak "github.com/isbm/bbpack-info"

type BBPakMarkdownFormat struct {
	packages []*bbpak.PackageMeta
}

func NewBBPakMarkdownFormat() *BBPakMarkdownFormat {
	bbp := new(BBPakMarkdownFormat)
	bbp.packages = make([]*bbpak.PackageMeta, 0)

	return bbp
}

// Format the output to the Markdown table (useful for GitHub/-Lab Wikis)
func (bbp *BBPakMarkdownFormat) Format() string {
	return ""
}

// SetPackages has been already collected and ready to format the output
func (bbp *BBPakMarkdownFormat) SetPackages(packages []*bbpak.PackageMeta) {
	bbp.packages = packages
}
