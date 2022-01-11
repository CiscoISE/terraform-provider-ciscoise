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

data "ciscoise_resource_version" "example" {
  provider = ciscoise
  resource = "networkdevice"
}

output "ciscoise_resource_version_example" {
  value = data.ciscoise_resource_version.example.item
}
