terraform {
  required_providers {
    ciscoise = {
      version = "0.0.1-beta"
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
  value = data.ciscoise_trusted_certificate.example.items
}

data "ciscoise_trusted_certificate" "example_id" {
  provider = ciscoise
  id       = data.ciscoise_trusted_certificate.example.items[0].id
}

output "ciscoise_trusted_certificate_example_id" {
  value = data.ciscoise_trusted_certificate.example_id.item
}
