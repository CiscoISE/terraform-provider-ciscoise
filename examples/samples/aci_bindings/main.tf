terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.4"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_aci_bindings" "response" {
  provider = ciscoise
}
output "ciscoise_aci_bindings_response" {
  value = data.ciscoise_aci_bindings.response
}
