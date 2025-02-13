terraform {
  required_providers {
    ciscoise = {
      version = "0.8.2-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_csr" "filtered" {
  provider = ciscoise
  filter   = ["friendlyName.EQ.ise#ISE Intermediate CA"]
  page     = 1
  size     = 1
  sort     = "asc"
  sort_by  = "timeStamp"
}

output "ciscoise_csr_filtered" {
  value = data.ciscoise_csr.filtered.items
}

data "ciscoise_csr" "example" {
  provider  = ciscoise
  host_name = data.ciscoise_csr.filtered.items[0].host_name
  id        = data.ciscoise_csr.filtered.items[0].id
}

output "ciscoise_csr_example" {
  sensitive = true
  value     = data.ciscoise_csr.example.item[0].csr_contents
}
