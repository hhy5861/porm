package dialect

import (
	"fmt"
	"strings"
	"time"
)

type (
	avatica struct {
	}
)

func (d avatica) QuoteIdent(s string) string {
	return s
}

func (d avatica) EncodeString(s string) string {
	var buf strings.Builder

	buf.WriteRune('\'')
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 0:
			buf.WriteString(`\0`)
		case '\'':
			buf.WriteString(`\'`)
		case '"':
			buf.WriteString(`\"`)
		case '\b':
			buf.WriteString(`\b`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case 26:
			buf.WriteString(`\Z`)
		case '\\':
			buf.WriteString(`\\`)
		default:
			buf.WriteByte(s[i])
		}
	}

	buf.WriteRune('\'')
	return buf.String()
}

func (d avatica) EncodeBool(b bool) string {
	if b {
		return "1"
	}

	return "0"
}

func (d avatica) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(timeFormat) + `'`
}

func (d avatica) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`0x%x`, b)
}

func (d avatica) Placeholder(_ int) string {
	return "?"
}
