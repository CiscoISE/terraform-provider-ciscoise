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

data "ciscoise_csr" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise__csr_response" {
  value = data.ciscoise_csr.response
}

# data "ciscoise_csr_export" "single_response" {
#   provider = ciscoise
#   hostname = "ise"
#   dirpath  = "/tmp/ise"
#   id       = data.ciscoise_csr.response.items[0].id
# }

# output "ciscoise__csr_export_single_response" {
#   value = data.ciscoise_csr_export.single_response
# }
