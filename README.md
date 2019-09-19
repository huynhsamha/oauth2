# Generic OAuth2 for Go

oauth2 package contains a client implementation for OAuth 2.0 spec.

Changes from https://github.com/golang/oauth2 to custom params.

## Installation

~~~~
go get github.com/huynhsamha/oauth2
~~~~

Or you can manually git clone the repository to
`$(go env GOPATH)/src/github.com/huynhsamha/oauth2`.

## Changes

+ In package `oauth2/internal`, add `params_conf.go`
+ Changes from hard code `client_id`, `client_secret` to params config
+ Add endpoint OAuth2 of Zalo Social API
+ Add `AuthStyleInQuery`, support *access_token* using HTTP GET
