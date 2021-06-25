<img align="right" width="100" src="https://kodo-cdn.dtapp.net/04/999e9f2f06d396968eacc10ce9bc8a.png" alt="www.dtapp.net"/>

<h1 align="left"><a href="https://www.dtapp.net/">Golang SSH Tunnel</a></h1>

📦 Golang SSH 隧道

通过Go实现从本地连接到远程服务器的数据库

[comment]: <> (dtapps)
![GitHub Org's stars](https://img.shields.io/github/stars/dtapps?style=for-the-badge)

[comment]: <> (go)
![golang version](https://img.shields.io/badge/golang-%3E%3D1.6-8892BF.svg?style=for-the-badge)
![godoc](https://pkg.go.dev/badge/github.com/dtapps/go-ssh-tunnel?status.svg)

[comment]: <> (goproxy.cn)
![goproxy](https://goproxy.cn/stats/github.com/dtapps/go-ssh-tunnel/badges/download-count.svg)

[comment]: <> (goreportcard.com)
![go report card](https://goreportcard.com/badge/github.com/dtapps/go-ssh-tunnel)

[comment]: <> (badge.fury.io)
![go project version](https://badge.fury.io/go/github.com%2Fdtapps%2Fgo-ssh-tunnel.svg)

[comment]: <> (github.com)
![latest release](https://badgen.net/github/release/dtapps/go-ssh-tunnel)
![latest stable release](https://badgen.net/github/release/dtapps/go-ssh-tunnel/stable)
![latest tag](https://badgen.net/github/tag/dtapps/go-ssh-tunnel)
![watchers](https://badgen.net/github/watchers/dtapps/go-ssh-tunnel)
![combined checks (master branch)](https://badgen.net/github/checks/dtapps/go-ssh-tunnel)
![stars](https://badgen.net/github/stars/dtapps/go-ssh-tunnel)
![forks](https://badgen.net/github/forks/dtapps/go-ssh-tunnel)
![issues](https://badgen.net/github/issues/dtapps/go-ssh-tunnel)
![branches](https://badgen.net/github/branches/dtapps/go-ssh-tunnel)
![releases](https://badgen.net/github/releases/dtapps/go-ssh-tunnel)
![tags](https://badgen.net/github/tags/dtapps/go-ssh-tunnel)
![license](https://badgen.net/github/license/dtapps/go-ssh-tunnel)
![contributors](https://badgen.net/github/contributors/dtapps/go-ssh-tunnel)
![assets downloads for latest release](https://badgen.net/github/assets-dl/dtapps/go-ssh-tunnel)
![assets downloads for a tag](https://badgen.net/github/assets-dl/dtapps/go-ssh-tunnel/1.0.11)
![repository dependents](https://badgen.net/github/dependents-repo/dtapps/go-ssh-tunnel)
![package dependents](https://badgen.net/github/dependents-pkg/dtapps/go-ssh-tunnel)
![dependabot status](https://badgen.net/github/dependabot/dtapps/go-ssh-tunnel)
![GitHub last commit](https://img.shields.io/github/last-commit/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/dtapps/go-ssh-tunnel?style=for-the-badge)
![Github All Contributors](https://img.shields.io/github/all-contributors/dtapps/go-ssh-tunnel/master?style=for-the-badge)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/dtapps/go-version/master?style=for-the-badge)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dtapps/go-ssh-tunnel/master?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub Sponsors](https://img.shields.io/github/sponsors/dtapps?style=for-the-badge)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub language count](https://img.shields.io/github/languages/count/dtapps/go-ssh-tunnel?style=for-the-badge)
![GitHub search hit counter](https://img.shields.io/github/search/dtapps/go-ssh-tunnel/go?style=for-the-badge)
![GitHub top language](https://img.shields.io/github/languages/top/dtapps/go-ssh-tunnel?style=for-the-badge)

[comment]: <> (sourcegraph.com)
![sourcegraph](https://sourcegraph.com/github.com/dtapps/go-ssh-tunnel/-/badge.svg)

[comment]: <> (scrutinizer-ci.com)
![Scrutinizer build (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/build/g/dtapps/go-ssh-tunnel/master?style=for-the-badge)
![Scrutinizer coverage (GitHub/BitBucket)](https://img.shields.io/scrutinizer/coverage/g/dtapps/go-ssh-tunnel/master?style=for-the-badge)
![Scrutinizer code quality (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/quality/g/dtapps/go-ssh-tunnel/master?style=for-the-badge)

[comment]: <> (www.codetriage.com)
![open source helpers](https://www.codetriage.com/dtapps/go-ssh-tunnel/badges/users.svg)

[comment]: <> (www.travis-ci.com)
![Travis (.com) branch](https://img.shields.io/travis/com/dtapps/go-ssh-tunnel/master?style=for-the-badge)

[comment]: <> (app.codecov.io)
![Codecov branch](https://img.shields.io/codecov/c/github/dtapps/go-ssh-tunnel/master?style=for-the-badge)

## 下载使用

- 把.example.config.yaml文件重命名为config.yaml
- 修改对应的配置，运行即可

## 扩展包使用

```go
go get -u github.com/dtapps/go-ssh-tunnel
```

```go
package main

import (
	"github.com/dtapps/go-ssh-tunnel"
)

func main() {
	Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
```