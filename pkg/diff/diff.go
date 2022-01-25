package diff

import "github.com/google/go-cmp/cmp"

func SliceDiff(a, b []string) string {
	return cmp.Diff(a, b)
}
