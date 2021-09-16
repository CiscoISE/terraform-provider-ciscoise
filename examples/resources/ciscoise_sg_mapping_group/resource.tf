
resource "ciscoise_sg_mapping_group" "example" {
    provider = ciscoise
    item {
      
      deploy_to = "string"
      deploy_type = "string"
      id = "string"
      name = "string"
      sgt = "string"
    }
}

output "ciscoise_sg_mapping_group_example" {
    value = ciscoise_sg_mapping_group.example
}