package engine

import (
	"path"
	"strings"
)

// ShiftPath shifts path one at a time to expose the current level in the path
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
