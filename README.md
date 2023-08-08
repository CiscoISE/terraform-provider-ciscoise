
# terraform-provider-ciscoise

terraform-provider-ciscoise is a Terraform Provider for [Cisco Identity Services Engine](https://developer.cisco.com/identity-services-engine/)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x
- [Go](https://golang.org/doc/install) 1.20 (to build the provider plugin)

## Introduction

The terraform-provider-ciscoise provides a Terraform provider for managing and automating your Cisco Identity Services Engine environment. It consists of a set of resources and data-sources for performing tasks related to Identity Services Engine.

This collection has been tested and supports Cisco Identity Services Engine 3.1_Patch_1.

## Using the provider

There are two ways to get and use the provider.
1. Downloading & installing it from registry.terraform.io
2. Building it from source

### From registry

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init. 

```hcl
terraform {
  required_providers {
    ciscoise = {
      source = "CiscoISE/ciscoise"
      version = "0.6.22-beta"
    }
  }
}

provider "ciscoise" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/CiscoISE/ciscoise/latest/docs#example-usage
}
```

### From build

Clone this repository to: `$GOPATH/src/github.com/CiscoISE/terraform-provider-ciscoise`

```sh
$ mkdir -p $GOPATH/src/github.com/CiscoISE/
$ cd $GOPATH/src/github.com/CiscoISE/
$ git clone https://github.com/CiscoISE/terraform-provider-ciscoise.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/CiscoISE/terraform-provider-ciscoise
$ make build
```

If the Makefile values (HOSTNAME, NAMESPACE, NAME, VERSION) were not changed, then the following code could used without changes.
Otherwise change the values accordingly.

To use this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    ciscoise = {
      source = "hashicorp.com/edu/ciscoise"
      version = "0.6.22-beta"
    }
  }
}

provider "ciscoise" {
  # Configuration options
  # More info at https://registry.terraform.io/providers/CiscoISE/ciscoise/latest/docs#example-usage
}
```


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.16+ is _required_). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-ciscoise
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources.

```sh
$ make testacc
```

## Documentation

In the docs directory you can find the documentation source code.

You can find the documentation online at [Terraform Registry - Cisco ISE provider](https://registry.terraform.io/providers/CiscoISE/ciscoise/latest/docs).

## Compatibility matrix
The following table shows the supported versions.

| Cisco ISE version | Terraform "ciscoise" provider version | Go "ciscoise-go-sdk" version|
|-------------------|---------------------------------------|---------------------|
| 3.1._Patch_1      | 0.6.22-beta                           | 1.1.19              |

If your SDK, Terraform provider is older please consider updating it first.

# Contributing

Ongoing development efforts and contributions to this provider are tracked as issues in this repository.

We welcome community contributions to this project. If you find problems, need an enhancement or need a new data-source or resource, please open an issue or create a PR against the [Terraform Provider for Cisco Identity Services Engine repository](https://github.com/CiscoISE/terraform-provider-ciscoise/issues).

# Change log

All notable changes to this project will be documented in the [CHANGELOG](./CHANGELOG.md) file.

The development team may make additional changes as the library evolves with the Cisco Identity Services Engine.

## License

This library is distributed under the license found in the [LICENSE](./LICENSE) file.
