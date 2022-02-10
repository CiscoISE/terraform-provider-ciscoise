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

data "ciscoise_aci_bindings" "response" {
  provider = ciscoise
}
output "ciscoise_aci_bindings_response" {
  value = data.ciscoise_aci_bindings.response
}
