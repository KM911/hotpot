hotpot  ♨️  let you eat hot tofu.

Watch file change and execute command.

[中文文档](https://github.com/KM911/hotpot/blob/main/README_zh-CN.md)

## Demo

### Restart Web Server

hotpot could help

![1701236594410](image/README/1701236594410.png)

## Install

1. download the binary from [github release](https://github.com/KM911/hotpot/releases)
2. build from source . Need go 1.18
3. go get ?

## Configuration

```toml
Delay = 2000
Command = 'go build'
WatchFiles = ['go'] # only watch xx.go files change 
IgnoreFolders = ['node_modules', 'vendor', '.git', '.idea', '.vscode', 'log', 'build', 'dist', 'bin', 'public', 'target', 'output']
ShowEvent = false
Github = 'https://github.com/KM911/hotpot'
```

## Acknowledgments

This project uses the following projects, for which I am grateful. Without their open source contributions, this project would not have been possible.

[GitHub - fsnotify/fsnotify: Cross-platform file system notifications for Go.](https://github.com/fsnotify/fsnotify)

[GitHub - pelletier/go-toml: Go library for the TOML file format](https://github.com/pelletier/go-toml)
