
resource "ciscoise_sg_mapping" "example" {
  provider = ciscoise
  item {

    deploy_to     = "string"
    deploy_type   = "string"
    host_ip       = "string"
    host_name     = "string"
    id            = "string"
    mapping_group = "string"
    name          = "string"
    sgt           = "string"
  }
}

output "ciscoise_sg_mapping_example" {
  value = ciscoise_sg_mapping.example
}