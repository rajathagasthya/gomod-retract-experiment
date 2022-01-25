package diff

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func SliceDiff(a, b []string) string {
	cmp.Diff(a, b, cmpopts.EquateEmpty(), cmpopts.SortSlices(func(x, y string) bool { return x < y }))
}
