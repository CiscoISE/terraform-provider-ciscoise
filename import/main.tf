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


resource "ciscoise_allowed_protocols" "sample" {}

resource "ciscoise_allowed_protocols" "sample2" {}

resource "ciscoise_allowed_protocols" "sample3" {}
