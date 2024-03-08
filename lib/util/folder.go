package util

import "os"

func IsFile(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(err)
	}
	return !fileInfo.IsDir()
}

func IsExist(file_ string) bool {
	_, err := os.Stat(file_)
	return err == nil
}

func IsDir(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(file_)
	}
	return fileInfo.IsDir()
}

func IsEmptyDir(src string) bool {
	files, err := os.ReadDir(src)
	if err != nil {
		panic(err)
	}
	return len(files) == 0
}

func ListDir(_src string) (srcNames []string, _dirNames []string) {
	files, err := os.ReadDir(_src)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			_dirNames = append(_dirNames, file.Name())
		} else {
			srcNames = append(srcNames, file.Name())
		}
	}
	return
}

// 可以写一个read deep dir 其实还是比较简单的不是面

func ListDirDeep(_src string) (srcs []string, _folders []string) {
	files, folders := ListDir(_src)
	srcs = append(srcs, files...)
	_folders = append(_folders, folders...)
	for _, folder := range folders {
		srcs_, _folders_ := ListDirDeep(folder)
		srcs = append(srcs, srcs_...)
		_folders = append(_folders, _folders_...)
	}
	return
}
