
resource "ciscoise_pan_ha" "example" {
  provider = ciscoise
  parameters {
    failed_attempts  = 1
    is_enabled       = "false"
    polling_interval = 1
    primary_health_check_node {
      hostname = "string"
    }
    secondary_health_check_node {
      hostname = "string"
    }
  }
}

output "ciscoise_pan_ha_example" {
  value = ciscoise_pan_ha.example
}
