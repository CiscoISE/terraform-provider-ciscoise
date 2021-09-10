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

data "ciscoise_aci_settings" "response" {
  provider = ciscoise
}
output "ciscoise__aci_settings_response" {
  value = data.ciscoise_aci_settings.response
}
