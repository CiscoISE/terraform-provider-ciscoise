terraform {
  required_providers {
    ciscoise = {
      version = "0.6.7-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_sg_mapping_group" "example" {
  provider = ciscoise
  parameters {

    deploy_to   = "70836740-8bff-11e6-996c-525400b48521"
    deploy_type = "NDG"
    # id          = "string"
    name = "Test_SG_group"
    sgt  = "944b2f30-8c01-11e6-996c-525400b48521"
  }
}

output "ciscoise_sg_mapping_group_example" {
  value = ciscoise_sg_mapping_group.example
}

data "ciscoise_sg_mapping_group" "examples" {
  provider = ciscoise

}

output "ciscoise_sg_mapping_group_examples" {
  value = data.ciscoise_sg_mapping_group.examples.items
}

data "ciscoise_sg_mapping_group" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sg_mapping_group.examples.items[0].id
}

output "ciscoise_sg_mapping_group_example1" {
  value = data.ciscoise_sg_mapping_group.example1.item
}
