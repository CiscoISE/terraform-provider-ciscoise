terraform {
  required_providers {
    ciscoise = {
      version = "0.6.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_aci_test_connectivity" "test_resp" {
  provider = ciscoise
}

output "test_resp" {
  value = data.ciscoise_aci_test_connectivity.test_resp.item[0].result
}
