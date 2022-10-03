terraform {
  required_providers {
    ciscoise = {
      version = "0.6.8-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


data "ciscoise_certificate_template" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_certificate_template_example" {
  value = data.ciscoise_certificate_template.example.items
}

# data "ciscoise_certificate_template" "item_name" {
#   provider = ciscoise
#   name     = data.ciscoise_certificate_template.example.items[0].name
# }

# output "ciscoise_certificate_template_item_name" {
#   value = data.ciscoise_certificate_template.item_name.item_name
# }

data "ciscoise_certificate_template" "item_id" {
  provider = ciscoise
  id       = data.ciscoise_certificate_template.example.items[0].id
}

output "ciscoise_certificate_template_item_id" {
  value = data.ciscoise_certificate_template.item_id.item_id
}
