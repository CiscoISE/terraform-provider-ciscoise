terraform {
  # Section defining the providers to be searched and installed (if needed)
  required_providers {
    ciscoise = {
      version = "0.6.7-beta"
      # This `source` is the local built version of the provider. 
      # It will change once the provider is publish to registry.terraform.io (to "CiscoISE/ciscoise")
      source = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
  # The user can set the arguments here or use environment variables
  # In this case I'm using environment variables
  # base_url
  # username
  # password
  # debug
  # ssl_verify
  # use_api_gateway
  # use_csrf_token
}

# Data sources: Query information from ISE
# Resources: Manage ISE resources (they perform create, update, delete operations)

# To reduce the number of data_sources created, the function executed depends on the parameters given

# The first example is going to get all telemetry info
data "ciscoise_telemetry_info" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise__telemetry_info_response" {
  value = data.ciscoise_telemetry_info.response
}

# The second example is going to get telemetry_info by id
data "ciscoise_telemetry_info" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_telemetry_info.response.items[0].id
}

output "ciscoise__telemetry_info_single_response" {
  value = data.ciscoise_telemetry_info.single_response
}
