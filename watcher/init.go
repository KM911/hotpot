package watcher

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/KM911/fish/adt"
	"github.com/KM911/fish/format"
	"github.com/KM911/fish/fs"
	"github.com/KM911/fish/system"
	"github.com/KM911/hotpot/config"
	"github.com/fsnotify/fsnotify"
)

var (
	IgnoreFolders = map[string]struct{}{}
	WatchFiles    = map[string]struct{}{}

	watcher  *fsnotify.Watcher
	err      error
	ok       bool
	Debounce func(func())
	event    fsnotify.Event
	cmd      *exec.Cmd
	hookCmds = make(chan *exec.Cmd, 10)
)

/*
Init Watch Arguments
*/
func InitWatcherArgs() {

	adt.SetAppend(WatchFiles, config.UserToml.WatchFiles)
	adt.SetAppend(IgnoreFolders, config.UserToml.IgnoreFolders)

	Debounce = system.NewDebounce(time.Duration(config.UserToml.Delay) * time.Millisecond)
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// filter ignore folders
	watchFolder := []string{}
	for _, folder := range ListAllDir() {
		if _, ok := IgnoreFolders[folder]; ok {
			continue
		}
		watchFolder = append(watchFolder, folder)
		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(config.UserToml.WatchFiles) == 1 && config.UserToml.WatchFiles[0] == "*" {
		EventHandle = EventHandleAll
	} else {
		EventHandle = EventHandleWithFileExtension
	}

	/*
		Display configuration information
	*/
	format.BlockMessage("watch folder", watchFolder)
	// format.NoteMessage("Ignores", strings.Join(config.UserToml.IgnoreFolders, ","))
	// format.InfoMessage("File Type", strings.Join(config.UserToml.WatchFiles, ","))
	format.Info("Execute Command :" + config.UserToml.ExecuteCommand)
	// format.InfoMessage("Delay", strconv.Itoa(config.UserToml.Delay))
	// format.InfoMessage("ShowEvent", strconv.FormatBool(config.UserToml.ShowEvent))

}

/*
return subdirectories of working directory
*/
func Folders() []string {
	var folders []string
	files, _ := os.ReadDir(fs.WorkingDirectory)
	folders = append(folders, fs.WorkingDirectory)
	for _, file := range files {
		if file.IsDir() {
			if _, ok = IgnoreFolders[file.Name()]; !ok {
				folders = append(folders, fs.WorkingDirectory+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return folders
}

func ListAllDir() []string {
	find := exec.Cmd{
		Path: "/usr/bin/fd",
		Args: []string{"", "-t", "d"},
	}
	out, err := find.Output()
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}
