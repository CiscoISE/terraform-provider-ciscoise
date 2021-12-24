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

data "ciscoise_csr_generate_intermediate_ca" "example" {
  provider = ciscoise
}
output "ciscoise_csr_generate_intermediate_ca_example" {
  value = data.ciscoise_csr_generate_intermediate_ca.example
}