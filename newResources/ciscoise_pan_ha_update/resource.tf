resource "ciscoise_pan_ha_update" "example" {
  provider         = ciscoise
  hostname         = "string"
  is_enabled       = "true"
  failed_attempts  = 5
  polling_interval = 120
  primary_health_check_node {
    hostname = "isenode"
  }
  secondary_health_check_node {
    hostname = "isenode"
  }
}