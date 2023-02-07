[![Go Report Card](https://goreportcard.com/badge/github.com/istio/client-go)](https://goreportcard.com/report/github.com/istio/client-go)
[![GolangCI](https://golangci.com/badges/github.com/istio/client-go.svg)](https://golangci.com/r/github.com/istio/client-go)
[![GoDoc](https://godoc.org/istio.io/client-go?status.svg)](https://godoc.org/istio.io/client-go)

# Golang Client Library for Istio APIs

This go library contains clients that talk to Istio resources in a Kubernetes cluster.

## How to get it?

If you are using go1.11+, you can simply use the following command to get the lastest client code:

```sh
go get istio.io/client-go@master
```

Or if you need to use a specific version of the client code, you can specify a specific version number. For example:

```sh
go get istio.io/client-go@release-1.4
```

The version number matches with official Istio versions for releases 1.4+.
