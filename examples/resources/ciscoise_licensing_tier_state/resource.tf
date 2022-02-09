resource "ciscoise_licensing_tier_state" "example" {
  provider = ciscoise
  parameters {
    name   = "string" # "ESSENTIAL", "ADVANTAGE", "PREMIER", "DEVICEADMIN"
    status = "string" # "ENABLED", "DISABLED"
  }
}

output "ciscoise_licensing_tier_state_example" {
  value = ciscoise_licensing_tier_state.example
}
