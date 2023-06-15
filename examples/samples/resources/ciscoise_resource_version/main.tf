terraform {
  required_providers {
    ciscoise = {
      version = "0.6.19-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_resource_version" "example" {
  provider = ciscoise
  resource = "networkdevice"
}

output "ciscoise_resource_version_example" {
  value = data.ciscoise_resource_version.example.item
}
