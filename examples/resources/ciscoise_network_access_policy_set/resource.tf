
resource "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
  parameters {

    condition {

      children {

        attribute_id     = "string"
        attribute_name   = "string"
        attribute_value  = "string"
        dictionary_name  = "string"
        dictionary_value = "string"
        end_date         = "string"
        name             = "string"
        operator         = "string"
        start_date       = "string"
      }
      condition_type = "string"
      name           = "string"
      id             = "string"
      description    = "string"
      is_negate      = "false"

    }
    default     = "false"
    description = "string"
    hit_counts  = 1
    id          = "string"
    is_proxy    = "false"

    name         = "string"
    rank         = 1
    service_name = "string"
    state        = "string"
  }
}

output "ciscoise_network_access_policy_set_example" {
  value = ciscoise_network_access_policy_set.example
}