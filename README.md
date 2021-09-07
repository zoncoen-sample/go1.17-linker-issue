# go1.17-linker-issue

Go 1.17 may have a linker bug.

```shell
fatal error: unreachable method called. linker bug?
```

## Steps to reproduce

```shell
$ go version
go version go1.17 linux/amd64

$ go build -buildmode=plugin -o plugin.so ./plugin

$ go run main.go
fatal error: unreachable method called. linker bug?

goroutine 1 [running]:
runtime.throw({0xa1aab8, 0xc00014e880})
...
exit status 2
```

### with scenarigo

```shell
$ go install github.com/zoncoen/scenarigo/cmd/scenarigo@v0.8.0

$ go build -buildmode=plugin -o plugin.so ./plugin

$ scenarigo run test.yaml
fatal error: unreachable method called. linker bug?

goroutine 21 [running]:
runtime.throw({0xcb5e67, 0xc000171980})
...
```
