
data "ciscoise_licensing_tier_state" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_tier_state_example" {
  value = data.ciscoise_licensing_tier_state.example.items
}
