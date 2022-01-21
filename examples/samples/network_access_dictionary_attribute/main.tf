terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.4"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}



data "ciscoise_network_access_dictionary_attribute" "example1" {
  provider        = ciscoise
  dictionary_name = "Test_dict"
  name            = "val2"
}

output "ciscoise_network_access_dictionary_attribute_example1" {
  value = data.ciscoise_network_access_dictionary_attribute.example1.item
}
