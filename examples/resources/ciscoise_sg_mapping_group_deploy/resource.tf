
resource "ciscoise_sg_mapping_group_deploy" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id = "string"
  }

}