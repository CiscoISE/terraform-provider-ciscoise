terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
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
