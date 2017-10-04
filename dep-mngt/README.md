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

#### Management the dependency 

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

## Govendor

### Installation

### Usage

## Godep

### Installation

### Usage


