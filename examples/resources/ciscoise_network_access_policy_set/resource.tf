
resource "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
  parameters {

    condition {

      attribute_name  = "string"
      attribute_value = "string"
      children {

        condition_type = "string"
        is_negate      = "false"

      }
      condition_type = "string"
      dates_range {

        end_date   = "string"
        start_date = "string"
      }
      dates_range_exception {

        end_date   = "string"
        start_date = "string"
      }
      description      = "string"
      dictionary_name  = "string"
      dictionary_value = "string"
      hours_range {

        end_time   = "string"
        start_time = "string"
      }
      hours_range_exception {

        end_time   = "string"
        start_time = "string"
      }
      id        = "string"
      is_negate = "false"

      name                = "string"
      operator            = "string"
      week_days           = ["string"]
      week_days_exception = ["string"]
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