
resource "ciscoise_px_grid_direct" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    additional_properties = {}
    attributes            = {}
    connector_name        = "string"
    connector_type        = "string"
    deltasync_schedule    = {}
    description           = "string"
    enabled               = "false"
    fullsync_schedule     = {}
    protocol              = "string"

    skip_certificate_validations = "false"
    url                          = {}

  }
}

output "ciscoise_px_grid_direct_example" {
  value = ciscoise_px_grid_direct.example
}