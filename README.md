<img width="100" src="https://www.dtapp.net/assets/img/logo.png" alt="https://www.dtapp.net/"/>

<h1><a href="https://www.dtapp.net/">Golang SSH Tunnel</a></h1>

📦 Golang SSH 隧道

[comment]: <> (dtapps)
[![GitHub Org's stars](https://img.shields.io/github/stars/dtapps)](https://github.com/dtapps)

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/go-ssh-tunnel?status.svg)](https://pkg.go.dev/github.com/dtapps/go-ssh-tunnel)
[![oproxy.cn](https://goproxy.cn/stats/github.com/dtapps/go-ssh-tunnel/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/go-ssh-tunnel)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/go-ssh-tunnel)](https://goreportcard.com/report/github.com/dtapps/go-ssh-tunnel)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgo-ssh-tunnel)

[comment]: <> (github.com)
[![watchers](https://badgen.net/github/watchers/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/watchers)
[![stars](https://badgen.net/github/stars/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/stargazers)
[![forks](https://badgen.net/github/forks/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/network/members)
[![issues](https://badgen.net/github/issues/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/issues)
[![branches](https://badgen.net/github/branches/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/branches)
[![releases](https://badgen.net/github/releases/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/releases)
[![tags](https://badgen.net/github/tags/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/tags)
[![license](https://badgen.net/github/license/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/blob/master/LICENSE)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/releases)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/tags)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/pulls)
[![GitHub issues](https://img.shields.io/github/issues/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/issues)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub language count](https://img.shields.io/github/languages/count/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub search hit counter](https://img.shields.io/github/search/dtapps/go-ssh-tunnel/go)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub top language](https://img.shields.io/github/languages/top/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)

## 下载使用

- 把.example.config.yaml文件重命名为config.yaml
- 修改对应的配置，运行即可

## 安装

```go
go get -u github.com/dtapps/dtapps/go-ssh-tunnel
```

```go
package main

import (
	"github.com/dtapps/go-ssh-tunnel/dssh"
)

func main() {
	dssh.Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
```