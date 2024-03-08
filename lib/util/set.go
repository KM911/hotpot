package util

// util.SetAppend(WatchFiles, config.UserToml.WatchFiles)
// util.SetAppend(IgnoreFolders, config.UserToml.IgnoreFolders)

func SetAppend(_m map[string]struct{}, _s []string) {
	for _, v := range _s {
		_m[v] = struct{}{}
	}
}
