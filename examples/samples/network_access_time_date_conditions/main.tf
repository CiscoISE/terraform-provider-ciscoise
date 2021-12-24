terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_time_date_conditions" "response" {
  provider = ciscoise
}
output "ciscoise__network_access_time_date_conditions_response" {
  value = data.ciscoise_network_access_time_date_conditions.response
}

data "ciscoise_network_access_time_date_conditions" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_network_access_time_date_conditions.response.items[0].id
}
output "ciscoise__network_access_time_date_conditions_single_response" {
  value = data.ciscoise_network_access_time_date_conditions.single_response
}
