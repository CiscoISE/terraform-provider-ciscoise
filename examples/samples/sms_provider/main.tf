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

data "ciscoise_sms_provider" "example" {
  provider = ciscoise

}

output "ciscoise_sms_provider_example" {
  value = data.ciscoise_sms_provider.example.items
}
