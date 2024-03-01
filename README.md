# Helper Functions

## Requirements

- [OpenTofu](https://opentofu.org) >= 1.x (Currently unknown)
- [Go](https://golang.org/doc/install) >= 1.21

## Building The Provider

1. Clone the repository
2. Enter the repository directory 
3. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

TODO:

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
export TF_ACC_TERRAFORM_PATH="/path/to/opentofu"
export TF_ACC_PROVIDER_NAMESPACE="hashicorp"
export TF_ACC_PROVIDER_HOST="registry.opentofu.org"
make testacc
```
