# Golang Dependency Management

## Dep

[Dep](https://github.com/golang/dep) is the official but still experimental Go dependency management tool. 
It is a new project which starts from Mar 2016, so it only supports Go 1.8 or 1.8+.

### Installation

Install with Homebrew on macOS:

```sh
$ brew install dep
$ brew upgrade dep
```

Install via `go get`:

```sh
go get -u github.com/golang/dep/cmd/dep
```

### Usage

#### Start the Management 

```sh
$ dep init
```

`dep init` will generate the manifest file `Gopkg.toml` and lock file `Gopkg.lock`, and install dependencies in 
`vendor/`.


#### Add dependencies

```sh
$ dep ensure
```

`dep ensure` will ensure the dependencies already in `vendor/` to match the constraints from the manifest, and install 
the latest version allowed by the manifest for the missing dependencies in `vendor/`. 

#### Add a dependency

```sh
$ dep ensure -add github.com/golang/glog
"github.com/golang/glog" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.
```

`dep ensure -add` will update `Gopkg.toml` and `Gopkg.lock`, and install the dependency in `vendor/`.
If the dependency has been imported in your code, just need to run `dep ensure`.

#### Update a dependency

1. Manually edit `Gopkg.toml`
2. Ensure the dependency

```sh
$ dep ensure
```

#### Remove a dependency

```sh
$ dep ensure
```

`dep ensure` will clean up the installed dependencies but not actually imported in code.

#### Check status of dependency

```sh
$ dep status
Lock inputs-digest mismatch due to the following packages missing from the lock:

PROJECT                   MISSING PACKAGES
github.com/docker/docker  [github.com/docker/docker/client]

This happens when a new import is added. Run `dep ensure` to install the missing packages.
```

After import new dependency in code or add directly add it into `Gopkg.toml`, `dep status` will find that they are 
missing in `Gopkg.lock`, and suggest you to run `dep ensure` to update lock file and install missing dependencies.

### Outputs

* **vendor**: Folder to store local dependencies(Has been renamed to `dep-vendor`).
* **Gopkg.toml**: Manifest file to describe the user intent for the dependencies.
* **Gopkg.lock**: Lock file to describe the status of the dependencies.

## Godep

[Godep](https://github.com/tools/godep) is one of the oldest dependency management tool for Go, which starts from Apr 2013. 
It can work with Go 1 or 1+.

### Installation

```sh
$ go get github.com/tools/godep
```

### Usage

#### Start the Management 

```sh
$ godep save ./...
```

`godep save` will generate the `Godeps/` folder and `Godeps/Godeps.json` file, and install dependencies in `vendor/`.

```sh
$ godep save github.com/supereagle/go-example/dep-mngt
godep: no buildable Go source files in /Users/robin/gocode/src/github.com/supereagle/go-example/dep-mngt
$ godep save github.com/supereagle/go-example/dep-mngt/cmd
```

`godep save` requires to specify the entrypoint(main package) of the application, it will capture all dependencies 
needed from this entrypoint. If your project has more than one entrypoints, you should run `godep save ./...` like 
`go test ./...`, `go install ./...` and `go fmt ./...`.

#### Add dependencies

```sh
$ godep save ./...
```

`godep save ./...` will copy the missing dependencies from `$GOPATH` to `vendor/`. 

#### Add a dependency

```sh
$ go get foo/bar
$ godep save ./...
```

Make sure the dependency to be added already is in `$GOPATH`, then `godep save ./...` will copy it from `$GOPATH` to 
`vendor/`.

#### Update a dependency

```sh
$ go get -u foo/bar
$ godep update foo/...
```

`godep update foo/...` will copy the latest dependency `foo` from `$GOPATH` to `vendor/`. 

#### Remove a dependency

**Not supported**. Need to remove the vendor and recreate it.

#### Check status of dependency

**Not supported**. No way to list the packages which are missing, or in `vendor/` but not used.

### Outputs

* **vendor**: Folder to store local dependencies(Has been renamed to `godep-vendor`).
* **Godeps/Godeps.json**: File to describe the status of the dependencies.

## Govendor

[Govendor](https://github.com/kardianos/govendor) is the vendor tool for Go. To learn more about it, please refer to its
[whitepaper](https://github.com/kardianos/govendor/blob/master/doc/whitepaper.md).

### Installation

```sh
go get -u github.com/kardianos/govendor
```

### Usage

#### Start the Management 

```sh
$ govendor init
$ govendor add +e
```

`govendor init` will create the `vendor/` folder and the `vendor/vendor.json` file.
`govendor add +e` will add existing files from `$GOPATH` to vendor.

#### Add dependencies

```sh
$ govendor add +e
```

`govendor add +e` will copy the missing dependencies from `$GOPATH` to `vendor/`. 

#### Add a dependency

```sh
$ govendor add -n github.com/golang/glog
  Copy "/Users/robin/gocode/src/github.com/golang/glog" -> 
  "/Users/robin/gocode/src/github.com/supereagle/go-example/dep-mngt/vendor/github.com/golang/glog"
	Ignore "glog_test.go"
```

`govendor add` will copy the packages from `$GOPATH` to `vendor/`.

> Option `-n`: dry run and print actions that would be taken.

#### Update a dependency

```sh
$ govendor update -n github.com/spf13/pflag
  Copy "/Users/robin/gocode/src/github.com/spf13/pflag" -> "/Users/robin/gocode/src/github.com/supereagle/go-example/dep-mngt/vendor/github.com/spf13/pflag"
  	Ignore "bool_slice_test.go"
  	Ignore "bool_test.go"
  	Ignore "count_test.go"
  	Ignore "example_test.go"
  	Ignore "export_test.go"
  	Ignore "flag_test.go"
  	Ignore "golangflag_test.go"
  	Ignore "int_slice_test.go"
  	Ignore "ip_slice_test.go"
  	Ignore "ip_test.go"
  	Ignore "ipnet_test.go"
  	Ignore "string_array_test.go"
  	Ignore "string_slice_test.go"
  	Ignore "uint_slice_test.go"
```

`govendor update` will update packages from `$GOPATH`. So you need to first make sure the correct versions of 
dependencies in `$GOPATH`. If the dependency is missing, it will be copied from `$GOPATH` to `vendor/`.

#### Remove a dependency

```sh
$ govendor remove -n  github.com/golang/glog
  Remove "/Users/robin/gocode/src/github.com/supereagle/go-example/dep-mngt/vendor/github.com/golang/glog/"
```

`govendor remove` will remove packages from the vendor folder.

#### Check status of dependency

```sh
$ govendor status
```

`govendor status` can not work correctly, as the missing packages can not be listed after `govendor remove` them.

### Outputs

* **vendor**: Folder to store local dependencies(Has been renamed to `govendor-vendor`).
* **vendor/vendor.json**: File to describe the status of the dependencies.
