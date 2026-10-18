package randutil

import "math/rand/v2"

func Choice[T any](slc []T) (T, bool) {
	if len(slc) < 1 {
		var n T
		return n, false
	}
	return slc[rand.IntN(len(slc))], true
}