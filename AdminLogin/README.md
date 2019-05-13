# AdminLogin Service

This is the AdminLogin service

Generated with

```
micro new MyShop/AdminLogin --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.AdminLogin
- Type: srv
- Alias: AdminLogin

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./AdminLogin-srv
```

Build a docker image
```
make docker
```