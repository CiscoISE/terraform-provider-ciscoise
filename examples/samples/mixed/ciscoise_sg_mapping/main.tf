terraform {
  required_providers {
    ciscoise = {
      version = "0.8.1-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_sg_mapping" "name" {
  provider = ciscoise
  parameters {

    deploy_to   = "70836740-8bff-11e6-996c-525400b48521"
    deploy_type = "NDG"
    host_ip     = "10.20.10.1/32"
    host_name   = ""
    # id            = "string"
    # mapping_group = ""
    name = "10.20.10.1/32"
    sgt  = "944b2f30-8c01-11e6-996c-525400b48521"
  }
}


output "ciscoise_sg_mapping_name" {
  value = ciscoise_sg_mapping.name
}

data "ciscoise_sg_mapping" "example" {
  provider = ciscoise

  depends_on = [
    ciscoise_sg_mapping.name
  ]
}

output "ciscoise_sg_mapping_example" {
  value = data.ciscoise_sg_mapping.example.items
}

data "ciscoise_sg_mapping" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sg_mapping.example.items[0].id
  depends_on = [
    ciscoise_sg_mapping.name
  ]
}

output "ciscoise_sg_mapping_example1" {
  value = data.ciscoise_sg_mapping.example1.item
}
