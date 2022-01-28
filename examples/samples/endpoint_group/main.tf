terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_endpoint_group" "examples" {
  provider = ciscoise
  count    = 2
  parameters {
    name        = "Sony-Device-${count.index}"
    description = "Identity Group for Profile: Sony-Device-0${count.index}"
  }
}

output "ciscoise_endpoint_group_examples" {
  value = ciscoise_endpoint_group.examples
}