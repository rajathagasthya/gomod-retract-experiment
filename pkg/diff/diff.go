package diff

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func SliceDiffIgnoreOrder(a, b []string) string {
	return cmp.Diff(a, b, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(x, y string) bool { return x < y }))
}
