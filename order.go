package porm

type direction bool

// orderby directions
// most databases by default use asc
const (
	asc  direction = false
	desc           = true
)

func order(column string, dir direction) Builder {
	return BuildFunc(func(d Dialect, buf Buffer) error {
		buf.WriteString(column)

		switch dir {
		case asc:
			buf.WriteString(" ")
		case desc:
			buf.WriteString(" ")
		}

		return nil
	})
}
