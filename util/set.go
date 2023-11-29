package util

func SetAppend(set map[string]struct{}, files []string) {
	for _, file := range files {
		set[file] = struct{}{}
	}
}
