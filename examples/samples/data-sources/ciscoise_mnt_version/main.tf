terraform {
  required_providers {
    ciscoise = {
      version = "0.6.13-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_version" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_version_example" {
  value = data.ciscoise_mnt_version.example.item
}
