
resource "ciscoise_upgrade_proceed" "example" {
  provider = ciscoise
  parameters {

    hostnames           = ["string"]
    pre_check_report_id = "string"
    upgrade_type        = "string"
  }
}

output "ciscoise_upgrade_proceed_example" {
  value = ciscoise_upgrade_proceed.example
}