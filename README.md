<img width="100" src="https://kodo-cdn.dtapp.net/04/999e9f2f06d396968eacc10ce9bc8a.png" alt="www.dtapp.net"/>

<h1><a href="https://www.dtapp.net/">Golang SSH Tunnel</a></h1>

📦 Golang SSH 隧道

通过Go实现从本地连接到远程服务器的数据库

[comment]: <> (dtapps)
[![GitHub Org's stars](https://img.shields.io/github/stars/dtapps)](https://github.com/dtapps)

[comment]: <> (go)
[![golang version](https://img.shields.io/badge/golang-%3E%3D1.6-8892BF.svg)](https://pkg.go.dev/github.com/dtapps/go-ssh-tunnel)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/go-ssh-tunnel?status.svg)](https://pkg.go.dev/github.com/dtapps/go-ssh-tunnel)

[comment]: <> (goproxy.cn)
[![goproxy](https://goproxy.cn/stats/github.com/dtapps/go-ssh-tunnel/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/go-ssh-tunnel)

[comment]: <> (goreportcard.com)
[![go report card](https://goreportcard.com/badge/github.com/dtapps/go-ssh-tunnel)](https://goreportcard.com/report/github.com/dtapps/go-ssh-tunnel)

[comment]: <> (github.com)
[![watchers](https://badgen.net/github/watchers/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/watchers)
[![stars](https://badgen.net/github/stars/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/stargazers)
[![forks](https://badgen.net/github/forks/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/network/members)
[![issues](https://badgen.net/github/issues/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/issues)
[![branches](https://badgen.net/github/branches/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/branches)
[![releases](https://badgen.net/github/releases/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/releases)
[![tags](https://badgen.net/github/tags/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/tags)
[![license](https://badgen.net/github/license/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/blob/master/LICENSE)
[![contributors](https://badgen.net/github/contributors/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/CONTRIBUTING.md)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/releases)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/tags)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/pulls)
[![GitHub issues](https://img.shields.io/github/issues/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel/issues)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/dtapps)](https://github.com/dtapps/go-ssh-tunnel/FUNDING.yml)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub language count](https://img.shields.io/github/languages/count/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub search hit counter](https://img.shields.io/github/search/dtapps/go-ssh-tunnel/go)](https://github.com/dtapps/go-ssh-tunnel)
[![GitHub top language](https://img.shields.io/github/languages/top/dtapps/go-ssh-tunnel)](https://github.com/dtapps/go-ssh-tunnel)

[comment]: <> (scrutinizer-ci.com)
[![Scrutinizer build (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/build/g/dtapps/go-ssh-tunnel/master)](https://scrutinizer-ci.com/g/dtapps/go-ssh-tunnel)
[![Scrutinizer coverage (GitHub/BitBucket)](https://img.shields.io/scrutinizer/coverage/g/dtapps/go-ssh-tunnel/master)](https://scrutinizer-ci.com/g/dtapps/go-ssh-tunnel)
[![Scrutinizer code quality (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/quality/g/dtapps/go-ssh-tunnel/master)](https://scrutinizer-ci.com/g/dtapps/go-ssh-tunnel)

[comment]: <> (www.travis-ci.com)
[![Travis (.com) branch](https://img.shields.io/travis/com/dtapps/go-ssh-tunnel/master)](https://www.travis-ci.com/github/dtapps/go-ssh-tunnel)

[comment]: <> (app.codecov.io)
[![Codecov branch](https://img.shields.io/codecov/c/github/dtapps/go-ssh-tunnel/master)](https://app.codecov.io/gh/dtapps/go-ssh-tunnel)

[comment]: <> (gitlab.com)
[![gitlab (.com)](https://gitlab.com/dtapps/go-ssh-tunnel/badges/master/pipeline.svg)](https://gitlab.com/dtapps/go-ssh-tunnel)

## 下载使用

- 把.example.config.yaml文件重命名为config.yaml
- 修改对应的配置，运行即可

## 扩展包使用

```go
go get -u github.com/dtapps/dtapps/go-ssh-tunnel
```

```go
package main

import (
	"github.com/dtapps/dtapps/go-ssh-tunnel"
)

func main() {
	Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
```