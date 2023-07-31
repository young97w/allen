package allen

import "strings"

type builder struct {
	sb     strings.Builder
	args   []any
	quoter byte
}
