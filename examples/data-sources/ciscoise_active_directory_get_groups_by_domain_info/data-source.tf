terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}
data "ciscoise_active_directory_get_groups_by_domain_info" "example" {
  provider = ciscoise
  id       = "4964ba10-4f1c-11ed-9aa4-6e36de26f9f6"
  additional_data {

    name  = "domain"
    value = "dcloud12121.com"
  }
}