
resource "ciscoise_sg_to_vn_to_vlan" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    sgt_id      = "string"
    virtualnetworklist {

      default_virtual_network = "false"
      description             = "string"
      id                      = "string"
      name                    = "string"
      vlans {

        data         = "false"
        default_vlan = "false"
        description  = "string"
        id           = "string"
        max_value    = 1
        name         = "string"
      }
    }
  }
}

output "ciscoise_sg_to_vn_to_vlan_example" {
  value = ciscoise_sg_to_vn_to_vlan.example
}