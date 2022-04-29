
resource "ciscoise_sg_mapping_deploy" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id = "string"
  }

}