package nue

import "strings"

func splitURLPath(path string) (prefix, pattern string) {
	path = path[1:]
	i := strings.Index(path, "/")
	if i < 0 {
		return "/" + path, ""
	}
	return "/" + path[:i], path[i:]
}
