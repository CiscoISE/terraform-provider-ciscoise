terraform {
  required_providers {
    ciscoise = {
      version = "0.4.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_portal_theme" "example" {
  provider = ciscoise

}

output "ciscoise_portal_theme_example" {
  value = data.ciscoise_portal_theme.example.items
}

data "ciscoise_portal_theme" "example1" {
  provider = ciscoise
  id       = data.ciscoise_portal_theme.example.items[0].id
}

output "ciscoise_portal_theme_example1" {
  value = data.ciscoise_portal_theme.example1.item
}
