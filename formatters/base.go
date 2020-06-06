package bbpak_formatters

import "fmt"

type BBPakFormatterUtils struct{}

func (bbp *BBPakFormatterUtils) ByteIEC(b uint64) string {
	const unit = 0x400
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
