hotpot  ♨️  让你吃上热豆腐.

监听文件修改并执行指定命令.

## 演示示例

### 重启web服务器

![img](image/README/1701236594410.png)

## 安装

1. 从 [github releases](https://github.com/KM911/hotpot/releases) 安装二进制文件.
2. go install

```bash
go install github.com/KM911/hotpot@latest
```

3. 从源码编译安装. 需要 go 1.18


## 实现原理

一句话概括就是 inotify + exec .

更加详细的原理解释可以看这里 [热编译技术原理](https://km911.github.io/obsidian-web-export/CODE/GO/%E7%83%AD%E7%BC%96%E8%AF%91%E6%8A%80%E6%9C%AF%E5%8E%9F%E7%90%86.html)

## 致谢

本项目使用了以下项目,对此他们表示感谢. 没有他们的开源行为,我不可能完成该项目.

[GitHub - fsnotify/fsnotify: Cross-platform file system notifications for Go.](https://github.com/fsnotify/fsnotify)
[GitHub - pelletier/go-toml: Go library for the TOML file format](https://github.com/pelletier/go-toml)
