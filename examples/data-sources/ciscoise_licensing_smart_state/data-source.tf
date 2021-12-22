
data "ciscoise_licensing_smart_state" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_smart_state_example" {
  value = data.ciscoise_licensing_smart_state.example.item
}
