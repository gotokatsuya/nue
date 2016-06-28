package nue

func splitURLPath(path string) (prefix, pattern string) {
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}
