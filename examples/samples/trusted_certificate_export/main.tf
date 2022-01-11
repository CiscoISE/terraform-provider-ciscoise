terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.2"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_trusted_certificate" "example" {
  provider = ciscoise
}

output "ciscoise_trusted_certificate_example" {
  value = data.ciscoise_trusted_certificate.example.items[0]
}

data "ciscoise_trusted_certificate_export" "example" {
  provider = ciscoise
  dirpath  = "/tmp/ise"
  id       = data.ciscoise_trusted_certificate.example.items[0].id
}
