terraform {
  required_providers {
    ciscoise = {
      version = "0.6.18-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

# It takes a couple of seconds/minutes to complete the pull deployment operation
data "ciscoise_deployment" "example" {
  provider = ciscoise
}

output "ciscoise_deployment_example" {
  value = data.ciscoise_deployment.example.item
}
