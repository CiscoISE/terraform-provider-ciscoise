terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_sg_mapping_deploy_all" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    
  }
}