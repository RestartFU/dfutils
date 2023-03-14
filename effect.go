package dfutil

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/df-mc/dragonfly/server/entity/effect"
)

// EffectName returns the name of the effect.
func EffectName(e effect.Effect) string {
	var s strings.Builder

	t := reflect.TypeOf(e.Type())
	name := t.Name()

	for _, r := range name {
		if unicode.IsUpper(r) && !strings.HasPrefix(name, string(r)) {
			s.WriteRune(' ')
		}
		s.WriteRune(r)
	}
	return s.String()
}
