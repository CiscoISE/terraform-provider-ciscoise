terraform {
  required_providers {
    ciscoise = {
      version = "1.0.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ise_telemetry_info" "response" {
  provider = ciscoise
  id       = "3753fbb4-e5c5-4a47-b5f6-3540adeac11d"
}
output "ise_telemetry_info_response" {
  value = data.ise_telemetry_info.response
}
