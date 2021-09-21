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

resource "ciscoise_endpoint_group" "examples" {
  provider = ciscoise
  count    = 2
  item {
    name        = "Sony-Device-${count.index}"
    description = "Identity Group for Profile: Sony-Device-${count.index}"
  }
}

output "ciscoise_endpoint_group_examples" {
  value = ciscoise_endpoint_group.examples
}