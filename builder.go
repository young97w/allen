package allen

import (
	"github.com/young97w/allen/model"
	"strings"
)

type builder struct {
	sb     strings.Builder
	model  *model.Model
	args   []any
	quoter byte
}
