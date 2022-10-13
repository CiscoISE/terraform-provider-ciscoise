terraform {
  required_providers {
    ciscoise = {
      version = "0.6.9-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_dictionary" "Test_dict" {
  provider = ciscoise
  parameters {
    name                 = "Test_dict"
    description          = "test dictionary"
    version              = "1.0"
    dictionary_attr_type = "ENTITY_ATTR"
  }
}

resource "ciscoise_network_access_dictionary_attribute" "example" {
  provider = ciscoise
  depends_on = [
    ciscoise_network_access_dictionary.Test_dict
  ]
  parameters {
    # allowed_values {
    #   is_default = "false"
    #   key        = "string"
    #   value      = "string"
    # }
    data_type       = "INT"
    description     = "value 2"
    dictionary_name = "Test_dict"
    direction_type  = "BOTH"
    internal_name   = "val2"
    name            = "val2"
  }
}

output "ciscoise_network_access_dictionary_attribute_example" {
  value = ciscoise_network_access_dictionary_attribute.example
}
