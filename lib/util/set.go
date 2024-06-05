package util

func SetAppend(_m map[string]struct{}, _s []string) {
	for _, v := range _s {
		_m[v] = struct{}{}
	}
}
