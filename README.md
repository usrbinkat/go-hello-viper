# go-hello-viper

Build `go-hello-viper-M.m.P-I.x86_64.rpm`
and   `go-hello-viper_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

A simple "go" program.
The purpose of the repository is to show how to:

1. Use "viper" go package.
1. See the "precedence" of variable values.

### Invocation

```console
go-hello-viper
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-hello-viper"
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-viper.git
```

#### Download dependencies

```console
cd ${REPOSITORY_DIR}
make dependencies
```

### Build

#### Local build

```console
cd ${REPOSITORY_DIR}
make build-local
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${REPOSITORY_DIR}
make build
```

The results will be in the `.../target` directory.

### Test

#### Manual

```console
export GOHELLOVIPER_OS_KEY="From OS"
export GOHELLOVIPER_OS_KEY_ONLY="From OS"
export GOHELLOVIPER_FLAG_KEY="From OS"
export GOHELLOVIPER_SET_KEY="From OS"

go-hello-viper --flagKey "From command option" --flagKeyOnly "From command option" --setKey "From command option"
```

#### Make-based test

```console
cd ${REPOSITORY_DIR}
make test-local
```

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh go-hello-viper-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh go-hello-viper-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i go-hello-viper_M.m.P-I_amd64.deb
```

### Cleanup

```console
cd ${REPOSITORY_DIR}
make clean
```
