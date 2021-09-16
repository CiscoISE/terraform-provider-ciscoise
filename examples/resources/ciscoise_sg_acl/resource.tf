
resource "ciscoise_sg_acl" "example" {
  provider = ciscoise
  item {

    aclcontent       = "string"
    description      = "string"
    generation_id    = "string"
    id               = "string"
    ip_version       = "string"
    is_read_only     = false
    modelled_content = "string"
    name             = "string"
  }
}

output "ciscoise_sg_acl_example" {
  value = ciscoise_sg_acl.example
}