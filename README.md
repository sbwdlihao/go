# Go学习

## Intellij使用

### IDE没有包缺失提示
需要安装[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)工具
```
$go get golang.org/x/tools/cmd/goimports
```

## 安装包遇到的问题

### 无法安装sha3包
```
$go get golang.org/x/crypto/sha3
package golang.org/x/crypto/sha3: unrecognized import path "golang.org/x/crypto/sha3" (https fetch: Get https://golang.org/x/crypto/sha3?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
```
国内网络的问题，解决办法是
```
$cd $GOPATH/golang.org/x
$git clone https://github.com/golang/crypto.git
```
