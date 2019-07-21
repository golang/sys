package consolidate

import (
	"go/ast"
	"testing"
)

func TestKindsConst(t *testing.T) {
	for _, tc := range []struct {
		label string
		a, b  *ast.ValueSpec
		res   *ast.ValueSpec
	}{
		{},
	} {
		t.Run(tc.label, func(t *testing.T) {
		})
	}
}
