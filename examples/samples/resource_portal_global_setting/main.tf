terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


data "ciscoise_portal_global_setting" "example" {
  provider = ciscoise
}

resource "ciscoise_portal_global_setting" "example" {
  provider = ciscoise
  parameters {
    id            = data.ciscoise_portal_global_setting.example.items[0].id
    customization = "HTMLANDJAVASCRIPT"
  }
}

output "ciscoise_portal_global_setting_example" {
  value = ciscoise_portal_global_setting.example
}