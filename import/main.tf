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


resource "ciscoise_allowed_protocols" "sample" {}

resource "ciscoise_allowed_protocols" "sample2" {}

resource "ciscoise_allowed_protocols" "sample3" {}
