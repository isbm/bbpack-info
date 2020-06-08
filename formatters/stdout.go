package bbpak_formatters

import (
	"fmt"
	"sort"
	"strings"

	"github.com/isbm/go-asciitable"
	"github.com/isbm/go-deb"
	"github.com/logrusorgru/aurora"

	bbpak_paktype "github.com/isbm/bbpack-info/paktype"
)

type BBPakSTDOUTFormat struct {
	BBPakFormatterUtils
}

func NewBBPakSTDOUTFormat() *BBPakSTDOUTFormat {
	bbp := new(BBPakSTDOUTFormat)
	bbp.packages = make(map[string]*bbpak_paktype.PackageMeta)

	return bbp
}

// Format the output to the ASCII text (for CLI for example)
func (bbp *BBPakSTDOUTFormat) Format() string {
	if len(bbp.packages) < 1 {
		return "Seems these packages were never deployed."
	} else {
		if bbp.IsSummary() {
			return bbp.formatSummary()
		} else {
			return bbp.formatDetails()
		}
	}
}

func (bbp *BBPakSTDOUTFormat) licenseColor(val string) string {
	if strings.Contains(strings.ToLower(val), "gpl") {
		if strings.Contains(val, "3") {
			return aurora.Red(val).String()
		} else {
			return aurora.Yellow(val).String()
		}
	} else {
		return aurora.Green(val).String()
	}
}

func (bbp *BBPakSTDOUTFormat) AddInfo(table *asciitable.TableData, name string, value string) {
	if value != "" {
		table.AddRow(aurora.Blue(name).String(), value)
	}
}

func (bbp *BBPakSTDOUTFormat) toTable(data *asciitable.TableData, firstColRight bool) string {
	style := asciitable.NewBorderStyle(asciitable.BORDER_SINGLE_THIN, asciitable.BORDER_NONE).
		SetGridVisible(false).SetHeaderVisible(true).SetBorderVisible(false).SetHeaderStyle(asciitable.BORDER_SINGLE_THICK)

	var align int
	if firstColRight {
		align = asciitable.ALIGN_RIGHT
	} else {
		align = asciitable.ALIGN_LEFT
	}
	return asciitable.NewSimpleTable(data, style).SetCellPadding(1).SetTextWrap(false).
		SetColAlign(align, 0).Render()
}

func (bbp *BBPakSTDOUTFormat) formatPackageSummary(p *bbpak_paktype.PackageMeta) string {
	var out strings.Builder
	out.WriteString(aurora.BrightWhite("Package ").String() +
		aurora.Bold(aurora.BrightWhite(p.GetPackage().ControlFile().Package())).String() +
		aurora.BrightWhite(" summary").String() + "\n")

	table := asciitable.NewTableData()
	table.SetHeader("TITLE", "INFO")

	bbp.AddInfo(table, "Package", aurora.White(p.GetPackage().ControlFile().Package()).String())
	bbp.AddInfo(table, "Summary", p.GetPackage().ControlFile().Summary())
	//bbp.AddInfo(table, "Description", p.GetPackage().ControlFile().Description())
	bbp.AddInfo(table, "Licence", bbp.licenseColor(p.GetPackage().ControlFile().Licence()))
	bbp.AddInfo(table, "Version", p.GetPackage().ControlFile().Version())
	bbp.AddInfo(table, "Priority", p.GetPackage().ControlFile().Priority())
	bbp.AddInfo(table, "Section", p.GetPackage().ControlFile().Section())
	bbp.AddInfo(table, "Source", p.GetPackage().ControlFile().Source())
	bbp.AddInfo(table, "Original Maintainer", p.GetPackage().ControlFile().OriginalMaintainer())
	bbp.AddInfo(table, "Maintainer", p.GetPackage().ControlFile().Maintainer())
	bbp.AddInfo(table, "Download Size", bbp.ByteIEC(p.GetPackage().FileSize()))

	out.WriteString(bbp.toTable(table, true))
	return out.String()
}

func (bbp *BBPakSTDOUTFormat) toGrid(values ...[]string) [][]interface{} {
	grid := make([][]interface{}, 0)
	maxlen := 0
	for _, vlist := range values {
		vlistLen := len(vlist)
		if maxlen < vlistLen {
			maxlen = vlistLen
		}
	}

	for i := 0; i < maxlen; i++ {
		row := make([]interface{}, 0)
		for _, vlist := range values {
			var val string
			if len(vlist) > i {
				val = vlist[i]
			} else {
				val = "-"
			}
			row = append(row, val)
		}
		grid = append(grid, row)
	}
	return grid
}

func (bbp *BBPakSTDOUTFormat) formatDepsMap(p *bbpak_paktype.PackageMeta) string {
	var out strings.Builder

	out.WriteString(aurora.Bold(aurora.BrightWhite("\n\nDependency Map\n")).String())

	table := asciitable.NewTableData()
	table.SetHeader("PROVIDES", "DEPENDS", "CONFLICTS")

	provides := p.GetPackage().ControlFile().Provides()
	depends := p.GetPackage().ControlFile().Depends()
	conflicts := p.GetPackage().ControlFile().Conflicts()

	pdcSet := false
	for _, row := range bbp.toGrid(provides, depends, conflicts) {
		table.AddRow(row...)
		pdcSet = true
	}

	if pdcSet {
		out.WriteString(bbp.toTable(table, false))
	} else {
		out.WriteString("This package provides nothing, depends on nothing and has no conflicts.")
	}
	return out.String()
}

func (bbp *BBPakSTDOUTFormat) formatScripts(p *bbpak_paktype.PackageMeta) string {
	var out strings.Builder

	preinstall := strings.TrimSpace(p.GetPackage().PreInstallScript())
	if preinstall != "" {
		out.WriteString(aurora.Bold(aurora.BrightWhite("\n\n\nPre-install Script\n\n")).String())
		out.WriteString(aurora.Yellow(preinstall).String())
	}

	postinstall := strings.TrimSpace(p.GetPackage().PostInstallScript())
	if postinstall != "" {
		out.WriteString(aurora.Bold(aurora.BrightWhite("\n\n\nPost-install Script\n\n")).String())
		out.WriteString(aurora.Yellow(postinstall).String())
	}

	preun := strings.TrimSpace(p.GetPackage().PreUninstallScript())
	if preun != "" {
		out.WriteString(aurora.Bold(aurora.BrightWhite("\n\n\nPre-uninstall Script\n\n")).String())
		out.WriteString(aurora.Yellow(preun).String())
	}

	postun := strings.TrimSpace(p.GetPackage().PostUninstallScript())
	if postun != "" {
		out.WriteString(aurora.Bold(aurora.BrightWhite("\n\n\nPost-uninstall Script\n\n")).String())
		out.WriteString(aurora.Yellow(postun).String())
	}

	return out.String()
}

// ColorContentFile
func (bbp *BBPakSTDOUTFormat) colorContentFilename(f deb.FileInfo) (string, bool) {
	if f.IsDir() || strings.HasSuffix(f.Name(), "/") { // This is odd
		return aurora.Index(25, f.Name()).String(), false
	} else if f.Linkname() != "" {
		return aurora.Index(45, f.Name()).String(), true
	} else if strings.Contains(f.Name(), "/bin/") || strings.Contains(f.Name(), "/sbin/") || strings.Contains(f.Name(), ".so") {
		return aurora.Bold(aurora.BrightGreen(f.Name())).String(), false
	}

	return f.Name(), false
}

func (bbp *BBPakSTDOUTFormat) formatContentFiles(p *bbpak_paktype.PackageMeta) string {
	var out strings.Builder
	out.WriteString(aurora.Bold(aurora.BrightWhite("\n\n\nPackage Contents\n")).String())

	table := asciitable.NewTableData()
	table.SetHeader("MODE", "UID", "GID", "SIZE", "NAME")

	for _, f := range p.GetPackage().Files() {
		fname, islink := bbp.colorContentFilename(f)
		table.AddRow(f.Mode().String(), f.Owner(), f.Group(), bbp.ByteIEC(uint64(f.Size())), fname)
		if islink {
			table.AddRow(f.Mode().String(), f.Owner(), f.Group(), bbp.ByteIEC(uint64(f.Size())), "\u21B3 "+f.Linkname())
		}
	}

	out.WriteString(bbp.toTable(table, false))
	return out.String()
}

// Return details about a single package
func (bbp *BBPakSTDOUTFormat) formatDetails() string {
	var out strings.Builder

	for _, pkg := range bbp.packages {
		out.WriteString(bbp.formatPackageSummary(pkg))
		out.WriteString(bbp.formatDepsMap(pkg))
		out.WriteString(bbp.formatScripts(pkg))
		out.WriteString(bbp.formatContentFiles(pkg))
		break
	}

	return out.String()
}

func (bbp *BBPakSTDOUTFormat) formatSummary() string {
	var out strings.Builder

	for idx, pkgName := range bbp.pkgNameIndex {
		pkg := bbp.packages[pkgName]
		out.WriteString(fmt.Sprintf("%3d. %s\n", idx+1, pkg.GetPackage().ControlFile().Package()))

		deps := pkg.GetPackage().ControlFile().Depends()
		if len(deps) > 0 {
			sort.Strings(deps)
			if len(deps) > 2 {
				out.WriteString("     Depends:\n")
				for didx, f := range deps {
					out.WriteString(fmt.Sprintf("     %3d. %s\n", didx+1, f))
				}
			} else {
				out.WriteString(fmt.Sprintf("     Depends:  %s\n", deps[0]))
			}
		}

		provides := pkg.GetPackage().ControlFile().Provides()
		if len(provides) > 0 {
			sort.Strings(provides)
			if len(provides) > 2 {
				out.WriteString("     Provides:\n")
				for didx, f := range provides {
					out.WriteString(fmt.Sprintf("     %3d. %s\n", didx+1, f))
				}
			} else {
				out.WriteString(fmt.Sprintf("     Provides: %s\n", provides[0]))
			}
		}

		out.WriteString(fmt.Sprintf("\n     License: %s, Maintainer: %s, Size: %s\n",
			pkg.GetPackage().ControlFile().Licence(),
			pkg.GetPackage().ControlFile().Maintainer(),
			bbp.ByteIEC(pkg.GetPackage().FileSize())))
		out.WriteString("\n")
	}

	return out.String()
}
